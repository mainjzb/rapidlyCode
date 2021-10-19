import requests  # 导入requests包
from bs4 import BeautifulSoup
from allURL import allJob
import re
import shutil

for job in allJob:
    url = 'https://maplestory.fandom.com/wiki/' + job + '/Skills'

    strhtml = requests.get(url)  # Get方式获取网页数据
    soup = BeautifulSoup(strhtml.text, 'lxml')
    skills = soup.select('#mw-content-text > div > table > tbody > tr:nth-child(1) > th:nth-child(2)')

    rule = ""
    with open("skill_" + job + ".md", 'a') as f:
        for skill in skills:
            skillSoup = BeautifulSoup(str(skill), 'lxml')
            skillImgs = skillSoup.select('a.image > img')
            ##mw-content-text > div > table:nth-child(10) > tbody > tr:nth-child(1) > th:nth-child(2) > a:nth-child(2)
            if len(skillImgs) > 0:
                skillName = skillSoup.find_all("a")[-1].get_text()
                rule = str(skillName).replace('-', '\\-') + "=="
                # print(skillName)
                for img in skillImgs:
                    # imgURL = re.search('src="(.*?)"', str(img)).group(1)
                    imgName = re.search('data-image-name="(.*?)"', str(img)).group(1)
                    rule += '<img src="upload/mxd/' + job + '/' + imgName + '"/>'
                    # print(imgName, imgURL)

                rule += skillName + "\n"
                f.write(rule.replace('／', '-'))
