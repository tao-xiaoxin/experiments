print("使用说明：必须安装Chrome浏览器否则不能使用本脚本")
print("弹出浏览器窗口请登录您的账户，脚本会为您自动抓取cookie！")

from selenium import webdriver

print("使用说明：必须安装chrome浏览器否则不能使用本脚本")
print("弹出浏览器窗口请登录您的账户，脚本会为您自动抓取cookie！")
url = "https://plogin.m.jd.com/login/login"

drive = webdriver.Chrome()
drive.get(url)
input("成功登录后请按回车继续:")
drive.refresh()
cookie = drive.get_cookies()
try:
    pt_key = cookie[20]
    pt_pin = cookie[9]

    pt_key = pt_key['value']
    pt_pin = pt_pin['value']

    pt_key = 'pt_key' + '=' + pt_key + ';'
    pt_pin = 'pt_pin' + '=' + pt_pin + ';'
    print("\n抓取成功！您的cookie:", pt_key, pt_pin)

    pt_pinw = str(pt_pin)
    pt_keyw = str(pt_key)
    with open("./jd_cookie.txt", 'w') as f:
        f.write(pt_keyw)
        f.write(pt_pinw)
        f.close()
    input("\n若出现pt_key...内容表示成功！请打开程序创建的cookie.txt文件取出京东的cookie\n按回车退出程序！")
    drive.close()
except Exception as e:
    print(f"An error occurred: {str(e)}")