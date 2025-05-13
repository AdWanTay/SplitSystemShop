import json
import sys
from random import randint
from time import sleep as pause
import urllib.request

import undetected_chromedriver as uc
from bs4 import BeautifulSoup
from selenium.webdriver.common.by import By
from tqdm import tqdm


def get_spec_value(spec, title_text, return_bool=False, return_first_word=False):
    for span in spec:
        if title_text in span.get_text(strip=True):
            spec_value_div = span.find_parent('li').find('div', class_='product-characteristics__spec-value')
            if spec_value_div:
                text = spec_value_div.get_text(strip=True)
                if return_bool:
                    return text.lower() == 'есть'
                if return_first_word:
                    return text.split()[0]
                return text
    return None if not return_bool else False


def parse_characteristics_page(driver, url):
    """Парсит страницу товара по ссылке."""
    driver.get(url)
    pause(randint(5, 7))
    soup = BeautifulSoup(driver.page_source, 'lxml')

    img_tag = soup.find('img', class_='product-images-slider__main-img')
    image_name = ""
    if img_tag and img_tag.get('src'):
        img_url = img_tag['src']
        print(f"Ссылка на изображение: {img_url}")
        image_name = img_url.split('/')[-1]
        urllib.request.urlretrieve(img_url, image_name)
    else:
        print("Изображение не найдено")

    driver.get(url+"characteristics")
    pause(randint(5, 7))

    expand_button = driver.find_element(By.CLASS_NAME, 'product-characteristics__expand')
    expand_button.click()
    pause(4)

    soup = BeautifulSoup(driver.page_source, 'lxml')

    # Парсинг различных данных с страницы
    title = soup.find('h1', class_="title")
    brand = soup.find('li', class_="breadcrumb-list__item initial-breadcrumb initial-breadcrumb_manufacturer").find('a',
                                                                                                                    class_="ui-link ui-link_black").find(
        'span')

    product_type = \
        soup.find('div', class_="product-characteristics-content").find_all('div',
                                                                            class_="product-characteristics__group")[
            1].find_all('li', class_="product-characteristics__spec")[2].find('a', class_="ui-link ui-link_blue")

    price = soup.find('div', class_="product-buy__price")
    if price:
        # Если блок содержит элемент с классом "product-buy__prev" (старая цена), то игнорируем её
        old_price = price.find('span', class_="product-buy__prev")

        if old_price:
            # Если есть старая цена, берем только цену из основного блока
            price_text = price.get_text(strip=True).split()[0:2]  # Берем только первую часть текста (основную цену)
        else:
            # Если старой цены нет, берем весь текст
            price_text = price.get_text(strip=True)

        # Извлекаем только цифры и преобразуем в целое число
        price_value = ''.join(c for c in price_text if c.isdigit())
        price_value = int(price_value) * 100 if price_value else 0  # Преобразуем в число
    else:
        price_value = 0

    spec = soup.find_all('span', class_='product-characteristics__spec-title-content')
    inverter_value = get_spec_value(spec, 'Инвертор', return_bool=True)
    recommended_area_value = get_spec_value(spec, 'Рекомендуемая площадь', return_first_word=True)
    cooling_power_value = get_spec_value(spec, 'Мощность охлаждения', return_first_word=True)
    cooling_class_value = get_spec_value(spec, 'Класс энергопотребления (охлаждение)')
    heating_class_value = get_spec_value(spec, 'Класс энергопотребления (обогрев)')
    min_noise_value = get_spec_value(spec, 'Минимальный уровень шума внутреннего блока', return_first_word=True)
    max_noise_value = get_spec_value(spec, 'Максимальный уровень шума внутреннего блока', return_first_word=True)

    external_width_value = get_spec_value(spec, "Ширина внешнего блока", return_first_word=True)
    external_height_value = get_spec_value(spec, "Высота внешнего блока", return_first_word=True)
    external_depth_value = get_spec_value(spec, "Глубина внешнего блока", return_first_word=True)
    external_weight_value = get_spec_value(spec, "Вес внешнего блока", return_first_word=True)
    internal_width_value = get_spec_value(spec, "Ширина внутреннего блока", return_first_word=True)
    internal_height_value = get_spec_value(spec, "Высота внутреннего блока", return_first_word=True)
    internal_depth_value = get_spec_value(spec, "Глубина внутреннего блока", return_first_word=True)
    internal_weight_value = get_spec_value(spec, "Вес внутреннего блока", return_first_word=True)

    modes_value = []
    for span in spec:
        if "Основные режимы" in span.get_text(strip=True):
            spec_value_div = span.find_parent('li').find('div', class_='product-characteristics__spec-value')
            if spec_value_div:
                text = spec_value_div.get_text(strip=True)
                modes_value = text.split(',')

    product_data = {
        "title": title.text.strip() if title else "",
        "brand": brand.text.strip() if brand else "",
        "product_type": product_type.text.strip() if product_type else "",
        "price": price_value,
        "has_inverter": inverter_value,
        "recommended_area": recommended_area_value,
        "cooling_power": float(cooling_power_value) / 1000 if cooling_power_value else None,
        "cooling_class": cooling_class_value if cooling_class_value else None,
        "heating_class": heating_class_value if heating_class_value else None,
        "min_noise": float(min_noise_value) if min_noise_value else None,
        "max_noise": float(max_noise_value) if max_noise_value else None,
        "external_width": int(external_width_value) if external_width_value else None,
        "external_height": int(external_height_value) if external_height_value else None,
        "external_depth": int(external_depth_value) if external_depth_value else None,
        "external_weight": float(external_weight_value) if external_weight_value else None,
        "internal_width": int(internal_width_value) if internal_width_value else None,
        "internal_height": int(internal_height_value) if internal_height_value else None,
        "internal_depth": int(internal_depth_value) if internal_depth_value else None,
        "internal_weight": float(internal_weight_value) if internal_weight_value else None,
        "modes": [mode.strip() for mode in modes_value],
        "image_url": image_name
    }

    return product_data


def get_all_category_page_urls(driver, url_to_parse):
    """Получаем URL категории и парсим ссылки с неё."""
    page = 1
    url = url_to_parse.format(page=page)
    driver.get(url=url)
    pause(10)

    soup = BeautifulSoup(driver.page_source, 'lxml')

    # span_tags = soup.find_all('span')
    # for i in span_tags:
    #     if bool(str(i).find('data-role="items-count"') != -1):
    #         number_of_pages = [int(x) for x in str(i) if x.isdigit()]
    #
    # res = int(''.join(map(str, number_of_pages)))
    # pages_total = ((res // 18) + 1)
    pages_total = 3
    print(f'Всего в категории {pages_total} страницы')

    urls = []

    while True:
        page_urls = get_urls_from_page(driver)
        urls += page_urls

        if page >= pages_total:
            break

        page += 1
        url = url_to_parse.format(page=page)
        driver.get(url)
        pause(randint(5, 5))

    return urls


def get_urls_from_page(driver):
    """Собирает все ссылки на текущей странице."""
    soup = BeautifulSoup(driver.page_source, 'lxml')
    elements = soup.find_all('a', class_="catalog-product__name ui-link ui-link_black")
    return list(map(
        lambda element: 'https://www.dns-shop.ru' + element.get("href"),
        elements
    ))


def save_to_json(data, file_name="output.json"):
    """Сохраняет данные в формате JSON"""
    with open(file_name, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=4)

    print(f"Данные успешно сохранены в файл {file_name}")


def main():
    driver = uc.Chrome()
    urls_to_parse = [
        'https://www.dns-shop.ru/catalog/17a8d3a316404e77/kondicionery/?p={page}',
    ]

    urls = []
    for index, url in enumerate(urls_to_parse):
        print(f'Получение списка всех ссылок из {index+1} категории:')
        parsed_url = get_all_category_page_urls(driver, url)
        urls.append(parsed_url)

    print("Запись всех ссылок в файл url.txt:")
    with open('urls.txt', 'w') as file:
        for url in urls:
            for link in url:
                file.write(link + "\n")

    with open('urls.txt', 'r') as file:
        urls = list(map(lambda line: line.strip(), file.readlines()))
        print(urls)
        info_dump = []
        for url in tqdm(urls, ncols=70, unit='товаров',
                        colour='blue', file=sys.stdout):
            info_dump.append(parse_characteristics_page(driver, url))

    save_to_json(info_dump)


if __name__ == '__main__':
    main()
    print('=' * 20)
    print('Все готово!')
