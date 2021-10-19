import requests  # 导入requests包
from bs4 import BeautifulSoup
import re
import shutil
from allURL import allJob
import os

for job in allJob:
    url = 'https://maplestory.fandom.com/wiki/' + job + '/Skills'
    strhtml = requests.get(url)  # Get方式获取网页数据
    soup = BeautifulSoup(strhtml.text, 'lxml')
    skills = soup.select('#mw-content-text > div > table > tbody > tr:nth-child(1) > th:nth-child(2)')

    for skill in skills:
        skillSoup = BeautifulSoup(str(skill), 'lxml')
        skillImgs = skillSoup.select('a.image > img')
        ##mw-content-text > div > table:nth-child(10) > tbody > tr:nth-child(1) > th:nth-child(2) > a:nth-child(2)
        if len(skillImgs) > 0:
            skillName = skillSoup.find_all("a")[-1].get_text()
            print(skillName)
            for img in skillImgs:
                imgURL = re.search('src="(.*?)"', str(img)).group(1)
                imgName = re.search('data-image-name="(.*?)"', str(img)).group(1)
                # download imgs save to local file
                r = requests.get(imgURL, stream=True)
                r.raw.decode_content = True
                if not os.path.exists(job):
                    os.makedirs(job)

                fileName = (job + "/" + imgName).replace('／', '-')
                with open(fileName, 'wb') as f:
                    shutil.copyfileobj(r.raw, f)
                # print(imgName, imgURL)

    # imgUrl = re.search('src="(.*?)"', str(img)).group(1)
    # imgName = re.search('data-image-name="(.*?)"', str(img)).group(1)
    # print(imgName, imgUrl)
