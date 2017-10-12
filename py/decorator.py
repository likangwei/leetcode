__author__ = 'likangwei'
import time
def how_quick(func):
    def wrap(*args, **kwargs):
        now = time.time()
        rst = func(*args, **kwargs)
        print func.__name__, time.time() - now
        return rst
    return wrap

