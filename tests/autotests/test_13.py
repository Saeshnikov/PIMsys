import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def generate_invalid_company_name():
    return '1' * 33

def test_console_error_on_invalid_company():
    options = webdriver.ChromeOptions()
    options.add_argument('--headless=new')
    options.add_argument('--disable-gpu')
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')
    options.set_capability('goog:loggingPrefs', {'browser': 'ALL'})

    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service, options=options)

    try:
        driver.get("http://ui:80/")
        wait = WebDriverWait(driver, 100)

        already_account_btn = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Уже есть аккаунт')]"))
        )
        already_account_btn.click()

        email_field = wait.until(EC.visibility_of_element_located((By.CSS_SELECTOR, "input[name='email']")))
        email_field.send_keys("admin")

        password_field = wait.until(EC.visibility_of_element_located((By.CSS_SELECTOR, "input[name='password']")))
        password_field.send_keys("adminadmin")

        login_button = wait.until(EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Войти')]")))
        login_button.click()

        wait.until(EC.url_to_be("http://ui:80/shop"))

        driver.get_log('browser')
        driver.execute_script("console.clear()")
        time.sleep(1)

        company_input = wait.until(
            EC.visibility_of_element_located((By.XPATH, "//input[contains(@class, 'MuiInputBase-input')]"))
        )
        company_input.click()
        company_input.clear()
        company_input.send_keys(generate_invalid_company_name())

        add_button = wait.until(EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить компанию')]")))
        add_button.click()

        start_time = time.time()
        error_found = False
        while time.time() - start_time < 20 and not error_found:
            logs = driver.get_log('browser')
            error_found = any(
                "Ошибка добавления магазина" in log['message'] and
                "value too long" in log['message']
                for log in logs
            )
            time.sleep(0.5)

        print("Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_console_error_on_invalid_company()