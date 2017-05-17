package main

import (
	"encoding/json"
	"flag"
	"github.com/beeker1121/goque"
	"github.com/golang/glog"
	"github.com/liuzl/newsmth/crawler"
	"github.com/liuzl/newsmth/downloader"
	"github.com/liuzl/newsmth/parser"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Process(task *crawler.Task) ([]parser.ParsedTask, []parser.ParsedItem) {
	glog.Info("proess: ", task.Url, ", task: ", task.Conf)
	req := &downloader.RequestInfo{Url: task.Url, Platform: "mobile"}
	resp := downloader.Download(req)
	if resp.Error != nil {
		glog.Error(resp.Error)
		return nil, nil
	}
	retUrls, retItems, err := parser.Parse(resp.Text, resp.Url, task.Conf)
	if err != nil {
		glog.Error(err)
		return nil, nil
	}
	return retUrls, retItems
}

func startWorker(worker int, wg *sync.WaitGroup, exitCh chan int, queue *goque.Queue, crawlerConf *crawler.CrawlerConf) {
	defer wg.Done()
	glog.Info("start worker: ", worker)
	defer glog.Info("exit worker: ", worker)
	for {
		select {
		case <-exitCh:
			return
		default:
			glog.Info("Work on next task! worker: ", worker)
			time.Sleep(1 * time.Second)
			glog.Info(queue.Length())
			item, err := queue.Dequeue()
			if err != nil {
				glog.Error(err)
				time.Sleep(1 * time.Second)
				continue
			}
			var task crawler.Task
			err = item.ToObject(&task)
			if err != nil {
				glog.Error(err)
			}
			tasks, items := Process(&task)
			glog.Info(tasks)
			glog.Info(items)
			for _, t := range tasks {
				queue.EnqueueObject(crawler.Task{Url: t.Url, Conf: crawlerConf.ParseConfs[t.TaskType]})
				break
			}
		}
	}
}

func Stop(sigs chan os.Signal, exitCh chan int) {
	<-sigs
	glog.Info("receive stop signal")
	close(exitCh)
}

func main() {
	defer glog.Flush()
	defer glog.Info("crawler exit")

	workerCnt := flag.Int("wc", 1, "worker count")
	confFile := flag.String("conf", "./www.newsmth.net.json", "crawler conf file")
	runDir := flag.String("run", "./run", "working dir")
	flag.Parse()
	if *workerCnt > 1000 {
		glog.Fatal("worker count too large, no larger than 1000")
	}
	conf, err := ioutil.ReadFile(*confFile)
	if err != nil {
		glog.Fatal(err)
	}
	var crawlerConf crawler.CrawlerConf
	err = json.Unmarshal(conf, &crawlerConf)
	if err != nil {
		glog.Fatal(err)
	}
	valid, err := crawler.IsValidCrawlerConf(crawlerConf)
	if err != nil {
		glog.Fatal(err)
	}
	if !valid {
		glog.Fatal("crawler conf invalid")
	}
	queue, err := goque.OpenQueue(*runDir)
	if err != nil {
		glog.Fatal(err)
	}
	defer queue.Close()
	if queue.Length() == 0 {
		glog.Info("New round of crawl")
		for _, url := range crawlerConf.StartUrls {
			task := crawler.Task{Url: url, Conf: crawlerConf.ParseConfs["start"]}
			glog.Info(task)
			if _, err := queue.EnqueueObject(task); err != nil {
				glog.Fatal(err)
			}
		}
	}

	exitCh := make(chan int)
	sigs := make(chan os.Signal)
	var wg sync.WaitGroup
	for i := 0; i < *workerCnt; i++ {
		wg.Add(1)
		go startWorker(i, &wg, exitCh, queue, &crawlerConf)
	}
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go Stop(sigs, exitCh)
	wg.Wait()
}
