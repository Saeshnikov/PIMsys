import random
import string
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def generate_random_company_name():
    return ''.join(random.choices(string.ascii_lowercase, k=7))

def generate_random_string(length=7):
    return ''.join(random.choices(string.ascii_lowercase, k=10))

def test_successful_company_creation():
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

        random_company_name = generate_random_company_name()

        company_input = wait.until(
            EC.visibility_of_element_located((
                By.XPATH,
                "//label[contains(text(), 'Название компании')]/following-sibling::div//input"
            ))
        )
        company_input.click()
        company_input.clear()
        company_input.send_keys(random_company_name)

        random_description = generate_random_string()

        description_field = wait.until(
            EC.visibility_of_element_located((
                By.XPATH,
                "//label[contains(text(), 'Описание компании')]/following-sibling::div//textarea"
            ))
        )
        description_field.click()
        description_field.clear()
        description_field.send_keys(random_description)

        url_field = wait.until(
            EC.visibility_of_element_located((
                By.XPATH,
                "//label[contains(text(), 'URL')]/following-sibling::div//input"
            ))
        )
        url_field.click()
        url_field.clear()
        url_field.send_keys("testurl")

        add_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить компанию')]"))
        )
        add_button.click()

        WebDriverWait(driver, 15).until(
            EC.invisibility_of_element_located((By.XPATH, "//div[contains(@class, 'MuiCircularProgress-root')]"))
        )

        created_company = wait.until(
            EC.presence_of_element_located((
                By.XPATH,
                f"//div[contains(@class, 'MuiGrid-item')]//*[contains(text(), '{random_description}')]"
            )
            )
        )

        print(f"Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_successful_company_creation()