import os
import time
import unittest

from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.firefox.options import Options


class TestFrontend(unittest.TestCase):

    def setUp(self):
        options = Options()
        options.headless = True
        self.browser = webdriver.Firefox(options=options)
        self.uri = os.getenv('FRONTEND_URI', 'http://localhost:3000')
        self.addCleanup(self.browser.quit)

    def test_index_title(self):
        self.browser.get(self.uri)
        self.assertEqual(self.browser.title, 'Voting App')

    def test_index_vote_a(self):
        self.browser.get(self.uri)
        time.sleep(2)
        btn_a = self.browser.find_element(By.ID, 'btn-vote-a')
        btn_a.click()
        time.sleep(2)
        lbl_success = self.browser.find_element(By.ID, 'lbl-success')
        assert lbl_success.text == 'Success!'

    def test_index_vote_b(self):
        self.browser.get(self.uri)
        time.sleep(2)
        btn_b = self.browser.find_element(By.ID, 'btn-vote-b')
        btn_b.click()
        time.sleep(2)
        lbl_success = self.browser.find_element(By.ID, 'lbl-success')
        assert lbl_success.text == 'Success!'


if __name__ == '__main__':
    unittest.main()
