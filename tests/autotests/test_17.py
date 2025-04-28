import time
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from webdriver_manager.chrome import ChromeDriverManager

def test_successful_product_addition1():
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

        fields = [
            ("Название продукта", "TV", "input"),
            ("Цена", "100", "input"),
            ("Количество", "1000", "input"),
            ("Вес", "10", "input"),
            ("Производитель", "LG", "input"),
            ("Модель", "NZ2", "input"),
            ("Спецификации", "213", "input"),
            ("Год выпуска", "2025", "input"),
            ("Энергопотребление", "100", "input")
        ]

        for label, value, tag in fields:
            element = wait.until(
                EC.visibility_of_element_located((
                    By.XPATH,
                    f"//label[contains(., '{label}')]/following-sibling::div//{tag}"
                ))
            )
            element.clear()
            element.send_keys(value)

        add_button = wait.until(
            EC.element_to_be_clickable((By.XPATH, "//button[contains(., 'Добавить продукт')]"))
        )
        add_button.click()

        WebDriverWait(driver, 15).until(
            EC.invisibility_of_element_located((By.XPATH, "//div[contains(@class, 'MuiCircularProgress-root')]"))
        )

        wait.until(
            EC.presence_of_element_located((By.XPATH, "//*[contains(text(), 'TV')]"))
        )

        print("Тест пройден")

    except Exception as e:
        print(f"Тест не пройден")
        raise

    finally:
        driver.quit()

if __name__ == "__main__":
    test_successful_product_addition1()