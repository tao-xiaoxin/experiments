# -*- coding: UTF-8 -*-
import uuid
import oss2
import os
import bson
import datetime
import requests
from loguru import logger
from pathlib import Path

def download(url_path):
    req = requests.get(url_path)
    filename = url_path.split('/')[-1]
    if req.status_code != 200:
        print('文件下载异常')
        return False
    base_path = Path(__file__).parent.parent
    img_path = base_path.joinpath('./img').resolve()  # 获得文件夹的绝对路径
    if not img_path.exists():  # 日志文件夹不存在就新建
        img_path.mkdir()
    file_path = f"{img_path}/{datetime.datetime.now().strftime('%Y%m%d')}_{filename}"
    try:
        with open(file_path, 'wb') as f:
            f.write(req.content)
            f.close()
    except Exception as e:
        print(e)
    return file_path

class AliYunOSS:

    def __init__(self, app_key, app_secret, oss_url, bucket_name, file_url):
        auth = oss2.Auth(app_key, app_secret)
        self.bucket = oss2.Bucket(auth, oss_url, bucket_name)
        self.url_prefix = file_url

    def get_archive_date(self):
        '''
        生成归档文档日期
        :return:
        '''
        archive_date = datetime.datetime.now().strftime("%Y/%m/")
        return archive_date

    def make_unique_id(self):
        """
        生成唯一的id
        :return:
        """
        unique_id = bson.ObjectId()
        name = str(unique_id)
        return name

    def upload_image_async(self, upload_dto, need_logging=False):
        """
        同步方式上传图片
        :param upload_dto: UploadFileDto 对象
        :return:
        """
        headers = {
            'Content-Type': upload_dto.content_type,
            'x-oss-meta-width': str(upload_dto.width),
            'x-oss-meta-height': str(upload_dto.height),
        }

        # 生成唯一的filename
        name = self.make_unique_id()

        result = self.bucket.put_object(name, upload_dto.file_data, headers=headers)

        return name, self.get_full_url(name)

    def get_image_meta(self, name):
        """
        返回 meta 信息. 不会全部返回阿里云的信息, 只会返回一部分
        :param name:
        :return:
        """
        oss_meta = self.bucket.head_object(name).headers

        meta = {
            'width': oss_meta['x-oss-meta-width'],
            'height': oss_meta['x-oss-meta-height'],
            'content_type': oss_meta['content-type']
        }
        return meta

    def upload_file_async(self, file_data, content_type, need_logging=False):
        """
        同步方式上传文件
        :param upload_dto: UploadFileDto 对象
        :return:
        """
        headers = {
            'Content-Type': content_type,
        }

        # 生成唯一的filename
        name = self.make_unique_id()

        result = self.bucket.put_object(name, file_data, headers=headers)

        return name, self.get_full_url(name)

    def get_full_url(self, name):
        """
        返回图片的网络路径
        :param name:
        :return:
        """

        return self.url_prefix + "/" + name

    def upload_local_file_sync(self, file_path, content_type, need_logging=False):
        """
        同步方式上传文件
        """
        headers = {
            'Content-Type': content_type
        }
        # 生成唯一的filename
        file_ext = os.path.splitext(file_path)[1]
        name = catalog + self.get_archive_date() + self.make_unique_id() + file_ext.lower()

        # 必须以二进制的方式打开文件，因为需要知道文件包含的字节数。
        with open(file_path, 'rb') as fileobj:
            fileobj.seek(0)  # 0 表示从开始，直到文件结束。
            # Tell方法用于返回当前位置。
            current = fileobj.tell()
            result = self.bucket.put_object(name, fileobj, headers=headers)


        return name, self.get_full_url(name)

if __name__ == '__main__':
    import sys
    if len(sys.argv) < 2:
        print("调用错误, 图片格式不对")
        sys.exit(1)
    oss_url = ""

    # 上传至md/目录
    catalog = 'md/'
    app_key = ''
    app_secret = ""
    # 节点名称
    bucket_name = ''
    file_url = ""
    a_oss = AliYunOSS(app_key, app_secret, oss_url, bucket_name, file_url)
    url_list = []
    for i in range(1, len(sys.argv)):
        path = sys.argv[i]
        if path.startswith("http"):
            path = download(path)
        logger.info(path)
        name, url = a_oss.upload_local_file_sync(path, 'image/png', True)
        url_list.append(url)

    print("Upload Success:")
    for url in url_list:
        print(url)