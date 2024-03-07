package main
  2
  3import (
  4	"flag"
  5	"fmt"
  6	"time"
  7
  8	"k8s.io/klog/v2"
  9
 10	v1 "k8s.io/api/core/v1"
 11	"k8s.io/apimachinery/pkg/fields"
 12	"k8s.io/apimachinery/pkg/util/runtime"
 13	"k8s.io/apimachinery/pkg/util/wait"
 14	"k8s.io/client-go/kubernetes"
 15	"k8s.io/client-go/tools/cache"
 16	"k8s.io/client-go/tools/clientcmd"
 17	"k8s.io/client-go/util/workqueue"
 18)
 19
 20// Controller demonstrates how to implement a controller with client-go.
 21type Controller struct {
 22	indexer  cache.Indexer
 23	queue    workqueue.RateLimitingInterface
 24	informer cache.Controller
 25}
 26
 27// NewController creates a new Controller.
 28func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
 29	return &Controller{
 30		informer: informer,
 31		indexer:  indexer,
 32		queue:    queue,
 33	}
 34}
 35
 36func (c *Controller) processNextItem() bool {
 37	// Wait until there is a new item in the working queue
 38	key, quit := c.queue.Get()
 39	if quit {
 40		return false
 41	}
 42	// Tell the queue that we are done with processing this key. 
 43	defer c.queue.Done(key)
 44
 45	// Invoke the method containing the business logic
 46	err := c.syncToStdout(key.(string))
 47	// Handle the error if something went wrong during the execution of the business logic
 48	c.handleErr(err, key)
 49	return true
 50}
 51
 52// syncToStdout is the business logic of the controller. In this controller it simply prints
 53// information about the pod to stdout. In case an error happened, it has to simply return the error.
 54// The retry logic should not be part of the business logic.
 55func (c *Controller) syncToStdout(key string) error {
 56	obj, exists, err := c.indexer.GetByKey(key)
 57	if err != nil {
 58		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
 59		return err
 60	}
 61
 62	if !exists {
 63		// Below we will warm up our cache with a Pod, so that we will see a delete for one pod
 64		fmt.Printf("Pod %s does not exist anymore\n", key)
 65	} else {
 66		// Note that you also have to check the uid if you have a local controlled resource, which
 67		// is dependent on the actual instance, to detect that a Pod was recreated with the same name
 68		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
 69	}
 70	return nil
 71}
 72
 73// handleErr checks if an error happened and makes sure we will retry later.
 74func (c *Controller) handleErr(err error, key interface{}) {
 75	if err == nil {
 76		// Forget about the #AddRateLimited history of the key on every successful synchronization.
 77		// This ensures that future processing of updates for this key is not delayed because of
 78		// an outdated error history.
 79		c.queue.Forget(key)
 80		return
 81	}
 82
 83	// This controller retries 5 times if something goes wrong. After that, it stops trying.
 84	if c.queue.NumRequeues(key) < 5 {
 85		klog.Infof("Error syncing pod %v: %v", key, err)
 86
 87		// Re-enqueue the key rate limited. Based on the rate limiter on the
 88		// queue and the re-enqueue history, the key will be processed later again.
 89		c.queue.AddRateLimited(key)
 90		return
 91	}
 92
 93	c.queue.Forget(key)
 94	// Report to an external entity that, even after several retries, we could not successfully process this key
 95	runtime.HandleError(err)
 96	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
 97}
 98
 99// Run begins watching and syncing.
100func (c *Controller) Run(workers int, stopCh chan struct{}) {
101	defer runtime.HandleCrash()
102
103	// Let the workers stop when we are done
104	defer c.queue.ShutDown()
105	klog.Info("Starting Pod controller")
106
107	go c.informer.Run(stopCh)
108
109	// Wait for all involved caches to be synced, before processing items from the queue is started
110	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
111		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
112		return
113	}
114
115	for i := 0; i < workers; i++ {
116		go wait.Until(c.runWorker, time.Second, stopCh)
117	}
118
119	<-stopCh
120	klog.Info("Stopping Pod controller")
121}
122
123func (c *Controller) runWorker() {
124	for c.processNextItem() {
125	}
126}
127
128func main() {
129	var kubeconfig string
130	var master string
131
132	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
133	flag.StringVar(&master, "master", "", "master url")
134	flag.Parse()
135
136	// creates the connection
137	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
138	if err != nil {
139		klog.Fatal(err)
140	}
141
142	// creates the clientset
143	clientset, err := kubernetes.NewForConfig(config)
144	if err != nil {
145		klog.Fatal(err)
146	}
147
148	// create the pod watcher
149	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
150
151	// create the workqueue
152	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
153
154	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
155	// whenever the cache is updated, the pod key is added to the workqueue.
156	// Note that when we finally process the item from the workqueue, we might see a newer version
157	// of the Pod than the version which was responsible for triggering the update.
158	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
159		AddFunc: func(obj interface{}) {
160			key, err := cache.MetaNamespaceKeyFunc(obj)
161			if err == nil {
162				queue.Add(key)
163			}
164		},
165		UpdateFunc: func(old interface{}, new interface{}) {
166			key, err := cache.MetaNamespaceKeyFunc(new)
167			if err == nil {
168				queue.Add(key)
169			}
170		},
171		DeleteFunc: func(obj interface{}) {
172			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
173			// key function.
174			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
175			if err == nil {
176				queue.Add(key)
177			}
178		},
179	}, cache.Indexers{})
180
181	controller := NewController(queue, indexer, informer)
182
183	// Now let's start the controller
184	stop := make(chan struct{})
185	defer close(stop)
186	go controller.Run(1, stop)
187
188	// Wait forever
189	select {}
190}