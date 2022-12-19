import time
import os

print("waiting 5 seconds...")

open("tp", "a").close()
time.sleep(20)

print("end")

os.remove("tp")
