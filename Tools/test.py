#!/usr/bin/python3 
#coding=utf-8

import platform
import os
import sys
import winreg
'''
              _ _        _____        ____  _____ 
    /\       | (_)      / ____|      / __ \|_   _|
   /  \      | |_  __ _| (___  _   _| |  | | | |  
  / /\ \ _   | | |/ _` |\___ \| | | | |  | | | |  
 / ____ \ |__| | | (_| |____) | |_| | |__| |_| |_ 
/_/    \_\____/|_|\__,_|_____/ \__,_|\___\_\_____|
'''
def logo():
	print(r"    ___           __   _             _____             ____      ____")
	print(r"   /   |         / /  (_)  ____ _   / ___/   __  __   / __ \    /  _/")
	print(r"  / /| |    __  / /  / /  / __ `/   \__ \   / / / /  / / / /    / /  ")
	print(r" / ___ |   / /_/ /  / /  / /_/ /   ___/ /  / /_/ /  / /_/ /   _/ /   ")
	print(r"/_/  |_|   \____/  /_/   \__,_/   /____/   \__,_/   \___\_\  /___/   ")
	print()

def find(keypath, name):
	key = winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE, keypath)
	value,type1 = winreg.QueryValueEx(key, name)
	del type1
	winreg.CloseKey(key)
	return value

def findRoot(keypath, name):
	key = winreg.OpenKey(winreg.HKEY_CLASSES_ROOT, keypath)
	value,type1 = winreg.QueryValueEx(key, name)
	del type1
	winreg.CloseKey(key)
	return value

def rall(keypath):
	key = winreg.OpenKey(winreg.HKEY_LOCAL_MACHINE, keypath)
	value = winreg.QueryInfoKey(key)
	winreg.CloseKey(key)
	return value

def CheckDotNet4_5():
	try:
		value = find(r"SOFTWARE\Microsoft\NET Framework Setup\NDP\v4\Full","Version" )
	except EnvironmentError:
		print("没有安装运行库.Net 4.5以上版本")
	else:
		print(".NET Framework版本: " + value)

def WindowsVersion():
	value = find(r"SOFTWARE\Microsoft\Windows NT\CurrentVersion", "BuildLabEx")
	i = value.find('.', 0)
	ii = value.find('.', i+1)
	osVersion = value[:ii]
	#print(platform.platform(True))

	version = platform.version()
	##print(version)
	if version[0:4] == '5.1.':
		print("系统版本: WinXP (" + osVersion + ")")
	elif version[0:4] == '6.1.':
		print("系统版本: Win7 (" + osVersion + ")")
	elif version[0:4] == '6.2.':
		print("系统版本: Win8 (" + osVersion + ")")
	elif version[0:5] == '10.0.':
		osVersion = value[:i]
		UBR = find(r"SOFTWARE\Microsoft\Windows NT\CurrentVersion", "UBR")
		win10Version = find(r"SOFTWARE\Microsoft\Windows NT\CurrentVersion", "ReleaseId")
		print("系统版本: Win10 " + win10Version + "(" + osVersion + "." + str(UBR) + ")")


def CheckVC2015():
	try:
		value = find(r"SOFTWARE\Classes\Installer\Dependencies\{d992c12e-cab2-426f-bde3-fb8c53950b0d}","Version" )
	except EnvironmentError:
		print("没有安装运行库VC++2015和VC++2017版本")
		return False
	else:
		print("VC++2015版本: " + value)
		return True

def CheckVC2017_2015():
	try:
		value = findRoot(r"Installer\Dependencies\,,amd64,14.0,bundle", "Version")
	except EnvironmentError:
		# find 2015
		try:
			value = find(r"SOFTWARE\Classes\Installer\Dependencies\{d992c12e-cab2-426f-bde3-fb8c53950b0d}","Version" )
		except EnvironmentError:
			print("没有安装运行库VC++2015和VC++2017版本")
			return False
		else:
			print("VC++2015版本: " + value)
			return True
	else:
		print("VC++ 2017版本: " + value)
		return True

def main():
	logo()
	WindowsVersion()
	CheckDotNet4_5()
	CheckVC2017_2015()
	print('------------------')
	input("按回车关闭窗口")

	
if __name__=="__main__":
	main()