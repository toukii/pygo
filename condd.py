#coding=utf-8
import random
import numpy as np
import json

# 相关性分析
def condd(*A, **args):
    print A
    print args

    start = args["start"]
    step = args["step"]
    end = args["end"]

    x = (args['end'] - args['start'])/ args['step'] +1
    y = len(A[0])-1

    print('%d 行 %d 列 ' % (x, y))
    # 随机数

    K = 0
    C = []
    for i in xrange(start,end+1,step):
        print i
        A1 = A[:i]
        r = np.corrcoef(A1, rowvar=0)
        r1 = r[0][1:]
        C.append(r1.tolist())
        K += 1

    print C
    resp = json.dumps(C)

    # resp = np.random.randn(x,y)
    # print(resp.tolist())
    # return json.dumps(resp.tolist())

    return resp

# demo
def foo(*args, **kwargs):
    s = "args=%s kwds=%s" % (args,kwargs)
    print(args)
    print(kwargs)
    print json.dumps((args, kwargs))
    return args, kwargs

if __name__ == "__main__":
    # print(foo())
    # print(foo(a=3))
    # kw=dict(a=3)
    # print(foo(**kw))

    start = 2
    end = 8
    args = {"start":start,"step":2,"end":end}
    I = [[1.9, 3.9, 3.9, 3.9, 3.9, 3.9, 8.9, 9.9], [1.88, 3.69, 3.69, 3.69, 3.69, 3.69, 8.45, 9.36], [2.65, 4.59, 4.59, 4.59, 4.59, 4.59, 5.86, 7.56], [3.12, 4.89, 4.89, 4.89, 4.89, 4.89, 6.32, 8.52], [3.25, 4.56, 4.56, 4.56, 4.56, 4.56, 7.25, 9.25], [3.46, 4.82, 4.82, 4.82, 4.82, 4.82, 7.14, 8.89], [3.65, 4.15, 4.15, 4.15, 4.15, 4.15, 4.52, 7.99], [4.21, 4.85, 4.85, 4.85, 4.85, 4.85, 5.12, 7.65]]
    print condd(*I,**args)
