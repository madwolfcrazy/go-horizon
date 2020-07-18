#coding=utf-8
import sys, os, re

def replacePackageName(fileFullPath, defaultPackageStr ,tarPackageStr):
    fr = open(fileFullPath, 'r')
    fileOldContent = fr.read()
    fr.close()
    fw = open(fileFullPath, 'w')
    fileNewContent = fileOldContent.replace(defaultPackageStr, tarPackageStr)
    fw.write(fileNewContent)
    fw.close()

def findGoModuleName(filename):
    fr = open(filename, 'r')
    content = fr.read()
    fr.close()
    m = re.search(r"^module (.*)", content)
    if m:
        return m.group(1)
    return ""

def run():
    args = sys.argv
    if len(args) < 2:
        print('Useage: "python init.py sample/samplename"')
        exit()
    defaultPackageStr = findGoModuleName("./go.mod")
    if defaultPackageStr == "":
        print("Can not find go.mod file")
        exit()
    tarPackageStr = args[1]

    for root, dirs, files in os.walk("."):
        for filename in files:

            fileFullPath = os.path.join(root, filename)
            if re.search(r"\w+(\.go)|(\.mod)$", fileFullPath):
                replacePackageName(fileFullPath, defaultPackageStr, tarPackageStr)



if __name__ == "__main__":
    run()