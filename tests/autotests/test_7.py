from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def test_registration_with_existing_user():
    options = webdriver.ChromeOptions()
    options.add_argument('--headless=new')
    options.add_argument('--disable-gpu')
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')

    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service, options=options)

    try:
        driver.get(" http://ui:80/")
        wait = WebDriverWait(driver, 100)

        wait.until(EC.visibility_of_element_located(
            (By.XPATH, "//*[contains(text(), 'Начните управлять своими товарами')]")
        ))

        fio_field = wait.until(EC.visibility_of_element_located((By.NAME, "name")))
        fio_field.click()
        fio_field.send_keys("name")

        phone_field = wait.until(EC.visibility_of_element_located((By.NAME, "phone")))
        phone_field.click()
        phone_field.send_keys("79001112233")

        email_field = wait.until(EC.visibility_of_element_located((By.NAME, "email")))
        email_field.click()
        email_field.send_keys("email@yandex.ru")

        password_field = wait.until(EC.visibility_of_element_located((By.NAME, "password")))
        password_field.click()
        password_field.send_keys("password")

        register_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//*[contains(text(), 'Зарегистрироваться')]"))
        )
        register_button.click()

        error_message = WebDriverWait(driver, 5).until(
            EC.visibility_of_element_located((By.XPATH, "//*[contains(text(), 'user already exists')]"))
        )

        if error_message:
            print("Тест пройден")

    except Exception as e:
        print("Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_registration_with_existing_user()