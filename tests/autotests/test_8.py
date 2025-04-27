from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def test_successful_login():
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
            EC.element_to_be_clickable((By.XPATH, "//*[contains(text(), 'Уже есть аккаунт')]"))
        )
        already_account_btn.click()

        wait.until(EC.visibility_of_element_located(
            (By.XPATH, "//*[contains(text(), 'Авторизация')]")
        ))

        email_field = wait.until(EC.visibility_of_element_located((By.NAME, "email")))
        email_field.send_keys("admin")

        password_field = wait.until(EC.visibility_of_element_located((By.NAME, "password")))
        password_field.send_keys("adminadmin")

        login_button = wait.until(EC.element_to_be_clickable((By.XPATH, "//*[contains(text(), 'Войти')]")))
        login_button.click()

        WebDriverWait(driver, 5).until(EC.url_to_be("http://localhost:3000/shop"))
        print("Тест пройден")

    except Exception as e:
        print("Тест не пройден")
        raise

    finally:
        driver.quit()

# Запуск
if __name__ == "__main__":
    test_successful_login()