import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def test_successful_product_addition4():
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

        wait.until(EC.url_to_be("http://localhost:3000/shop"))

        branches_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//div[@role='button' and .//span[text()='Филиалы']]"))
        )
        branches_button.click()

        products_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//div[@role='button' and .//span[text()='Категории']]"))
        )
        products_button.click()

        category_name_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Название категории')]/following-sibling::div//input"))
        )
        category_name_input.send_keys("test")

        description_category_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                "//label[contains(text(), 'Описание категории')]/following-sibling::div//textarea"))
        )
        description_category_input.send_keys("testing")

        add_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., '+ Добавить аттрибут')]"))
        )
        add_button.click()

        atribute_name_input = wait.until(
            EC.visibility_of_element_located((By.XPATH,
                 "//label[contains(text(), 'Название атрибута')]/following-sibling::div//input"))
        )
        atribute_name_input.send_keys("test attribute")

        add_attribute_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить категорию')]"))
        )
        add_attribute_button.click()

        created_attribute = wait.until(
            EC.presence_of_element_located((By.XPATH,
                     "//div[contains(@class, 'MuiGrid-item')]//*[contains(text(), 'test')]"))
        )

        print("Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_successful_product_addition4()