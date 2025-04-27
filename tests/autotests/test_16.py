import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager
from selenium.common.exceptions import TimeoutException

def test_successful_branch_creation_and_deletion():
    options = webdriver.ChromeOptions()
    options.add_argument('--headless=new')
    options.add_argument('--disable-gpu')
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')

    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service, options=options)

    try:
        driver.get("http://ui:80/")
        wait = WebDriverWait(driver, 100)

        already_account_btn = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Уже есть аккаунт')]"))
        )
        already_account_btn.click()

        email_field = wait.until(
            EC.visibility_of_element_located((By.CSS_SELECTOR, "input[name='email']"))
        )
        email_field.send_keys("admin")

        password_field = wait.until(
            EC.visibility_of_element_located((By.CSS_SELECTOR, "input[name='password']"))
        )
        password_field.send_keys("adminadmin")

        login_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Войти')]"))
        )
        login_button.click()

        wait.until(EC.url_to_be("http://ui:80/shop"))

        branches_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//div[@role='button' and .//span[text()='Филиалы']]"))
        )
        branches_button.click()

        branch_name_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Название филиала')]/following-sibling::div//input"))
        )
        branch_name_input.send_keys("testfil")

        branch_description_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Описание филиала')]/following-sibling::div//textarea"))
        )
        branch_description_input.send_keys("testing")

        branch_address_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Адрес филиала')]/following-sibling::div//input"))
        )
        branch_address_input.send_keys("test")

        branch_website_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Сайт филиала')]/following-sibling::div//input"))
        )
        branch_website_input.send_keys("test.com")

        add_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить филиал')]"))
        )
        add_button.click()

        WebDriverWait(driver, 15).until(
            EC.invisibility_of_element_located((By.XPATH, "//div[contains(@class, 'MuiCircularProgress-root')]"))
        )

        delete_button = WebDriverWait(driver, 25).until(
            EC.element_to_be_clickable((
                By.XPATH,
                "//*[contains(text(), 'testfil')]/ancestor::div[contains(@class, 'MuiPaper-root')]"
                "//button[contains(@class, 'MuiIconButton-colorDanger')]"
            ))
        )
        delete_button.click()

        try:
            WebDriverWait(driver, 15).until(
                EC.invisibility_of_element_located((
                    By.XPATH,
                    "//*[contains(text(), 'testfil')]/ancestor::div[contains(@class, 'MuiGrid-item')]"
                ))
            )
        except TimeoutException:
            raise AssertionError("Филиал не был удален в течение 15 секунд")

        print("Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_successful_branch_creation_and_deletion()