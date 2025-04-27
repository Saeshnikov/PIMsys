import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager


def test_successful_product_addition2():
    options = webdriver.ChromeOptions()
    options.add_argument('--headless=new')
    options.add_argument('--disable-gpu')
    options.add_argument('--window-size=1920,1080')
    options.add_argument('--no-sandbox')
    options.add_argument('--disable-dev-shm-usage')

    service = Service(ChromeDriverManager().install())
    driver = webdriver.Chrome(service=service, options=options)

    try:
        driver.get("http://localhost:3000/")
        wait = WebDriverWait(driver, 20)

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
            EC.element_to_be_clickable((By.XPATH, "//div[@role='button' and .//span[text()='Продукты']]"))
        )
        products_button.click()

        category_combobox = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//div[@role='combobox' and contains(@class, 'MuiSelect-select')]"))
        )
        category_combobox.click()

        tech_category = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//li[contains(., 'Техника')]"))
        )
        tech_category.click()

        def safe_input(field_locator, value):
            field = wait.until(EC.element_to_be_clickable(field_locator))
            field.click()

            for _ in range(3):
                field.send_keys(Keys.BACKSPACE)
                field.send_keys(Keys.DELETE)

            field.send_keys(value)
            WebDriverWait(driver, 3).until(
                lambda d: field.get_attribute('value') == str(value)
            )
        safe_input(
            (By.XPATH, "//label[contains(., 'Название продукта')]/following-sibling::div//input"),
            "TV"
        )

        safe_input(
            (By.XPATH, "//label[contains(., 'Цена')]/following-sibling::div//input"),
            "100"
        )

        safe_input(
            (By.XPATH, "//label[contains(., 'Количество')]/following-sibling::div//input"),
            "-1"
        )

        add_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить продукт')]"))
        )
        add_button.click()

        error_message = WebDriverWait(driver, 5).until(
            EC.visibility_of_element_located((By.XPATH, "//*[contains(text(), 'amount must be positive or null')]"))
        )

        if error_message:
            print("Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_successful_product_addition2()