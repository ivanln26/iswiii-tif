import os
import unittest

from selenium import webdriver
from selenium.webdriver.firefox.options import Options


class TestFrontend(unittest.TestCase):

    def setUp(self):
        options = Options()
        options.headless = True
        self.browser = webdriver.Firefox(options=options)
        self.uri = os.getenv('FRONTEND_URI', 'http://localhost:3000')
        self.addCleanup(self.browser.quit)

    def test_index(self):
        self.browser.get(self.uri)
        self.assertEqual(self.browser.title, 'Voting App')


if __name__ == '__main__':
    unittest.main()
