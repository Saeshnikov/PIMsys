from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

options = webdriver.ChromeOptions()
options.add_argument('--headless=new')  # Новый формат headless
options.add_argument('--disable-gpu')
options.add_argument('--window-size=1920,1080')
options.add_argument('--no-sandbox')
options.add_argument('--disable-dev-shm-usage')

# Правильная инициализация драйвера
service = Service(ChromeDriverManager().install())
driver = webdriver.Chrome(service=service, options=options)

try:
    driver.get(" http://ui:80/")
    wait = WebDriverWait(driver, 100)

    already_account_btn = wait.until(
        EC.element_to_be_clickable((By.XPATH, "//*[contains(text(), 'Уже есть аккаунт')]"))
    )
    already_account_btn.click()

    wait.until(EC.visibility_of_element_located(
        (By.XPATH, "//*[contains(text(), 'Авторизация')]")
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
    print(f"Ошибка: {str(e)}")
    raise

finally:
    driver.quit()