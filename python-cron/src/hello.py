import datetime

with open('hello.txt','a') as outFile:
    outFile.write('\n' + str(datetime.datetime.now()))
