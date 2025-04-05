from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--disable-gpu')
options.add_argument('--window-size=1920,1080')

driver = webdriver.Chrome(options=options)

try:
    driver.get("http://localhost:3000/")
    wait = WebDriverWait(driver, 10)

    already_account_btn = wait.until(
        EC.element_to_be_clickable((By.XPATH, "//[contains(text(), 'Уже есть аккаунт')]"))
    )
    already_account_btn.click()

    wait.until(EC.visibility_of_element_located(
        (By.XPATH, "//[contains(text(), 'Авторизация')]")
    ))

    email_field = wait.until(
        EC.visibility_of_element_located((By.NAME, "email"))
    )
    email_field.send_keys("admin")

    password_field = wait.until(
        EC.visibility_of_element_located((By.NAME, "password"))
    )
    password_field.send_keys("adminadmin")

    login_button = wait.until(
        EC.element_to_be_clickable((By.XPATH, "//*[contains(text(), 'Войти')]"))
    )
    login_button.click()

    print("Тест пройден успешно!")

except Exception as e:
    print(f"Ошибка при выполнении теста: {e}")

finally:
    driver.quit()
