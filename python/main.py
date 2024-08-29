import queue
import threading

import os
from dotenv import load_dotenv
from os import walk

from utils.redis import *

from tools.pdf.pyPdf2 import process as pypdf2
from tools.pdf.pdfPlumber import process as pdfplumber

# processFile = 0
# doneFile = 0

load_dotenv()

class Task():
    def __init__(self): 
        self.Read = None
        self.Write = None

r = Redis(os.getenv('REDIS_CONN'))

# The queue for tasks
q = queue.Queue()

# Worker, handles each task
def worker():
    while True:
        item = q.get()
        if item is None:
            break
        print("Working on", item.Read)
        # processFile = processFile + 1
        #rawText = pypdf2(readFileName)
        rawText = pdfplumber(item.Read)

        if len(rawText) > 0:
            r.client().set(item.Write,rawText,ex=30)
            # # Write file
            # with open(item.Write, 'w', encoding='utf-8') as f:
            #     print("Write File : ",item.Write)
            #     f.write(rawText)
                # doneFile = doneFile + 1
        q.task_done()


def start_workers(worker_pool=1000):
    threads = []
    for i in range(worker_pool):
        t = threading.Thread(target=worker, daemon=True)
        t.start()
        threads.append(t)
    return threads


def stop_workers(threads):
    # stop workers
    for i in threads:
        q.put(None)
    for t in threads:
        t.join()


def create_queue(task_items):
    for item in task_items:
        q.put(item)

if __name__ == "__main__":

    # Set up Tasks
    tasks = []
    # Open directory
    files = []

    for (dirpath, dirnames, filenames) in walk("../assets/"):
        files.extend(filenames)
        break

    for file in files:
        if ".pdf" not in file: 
            continue

        readFileName = dirpath+file
        writeFileName = file.replace(".pdf", "_raw")

        # processFile = processFile + 1

        task = Task()
        task.Read = readFileName
        task.Write = writeFileName
        tasks.append(task)

    # Start up your workers
    workers = start_workers(worker_pool=100)
    create_queue(tasks)

    # Blocks until all tasks are complete
    q.join()

    stop_workers(workers)

# print("done",doneFile,"/",processFile)