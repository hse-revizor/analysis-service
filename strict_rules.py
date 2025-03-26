import json
import re
import sys

def check_has_string_in_file(file_content: str, substring: str) -> dict:
    """ Проверяет, содержится ли ОДНА строка в файле (игнорируя регистр и пробелы) """

    lines = [line.lower().replace(" ", "") for line in file_content.split("\n")]
    normalized_substring = substring.lower().replace(" ", "")

    found = any(normalized_substring in line for line in lines)

    return {
        "status": "passed" if found else "failed",
        "details": {
            "message": "Строка найдена в файле." if found else "Строка не найдена."
        }
    }


def check_has_expected_value_in_field(file_content: str, field_name: str, expected_value: str) -> dict:
    """
    Проверяет, содержит ли указанная переменная в файле ожидаемое значение.
    Поддерживает присваивание через `=`, `:`, `=>`, `<-`, `.set()`, `let/var/const`, `settings[...]`
    """

    normalized_content = file_content.lower()
    normalized_field_name = re.escape(field_name.lower().strip())
    normalized_expected_value = expected_value.lower().strip()

    pattern = rf"""
        (?:(?:let|var|const)\s+)?  # Опциональные ключевые слова (let, var, const)
        {normalized_field_name}\s*  # Имя переменной
        ([:=><-]+|\.\s*set\()\s*  # Разделитель
        ["']?([^"'\n]+)["']?  # Значение (в кавычках или без)
    """

    matches = re.findall(pattern, normalized_content, re.MULTILINE | re.VERBOSE)

    if not matches:
        return {
            "status": "failed",
            "details": {
                "message": f"Поле '{field_name}' не найдено в файле.",
                "field": field_name
            }
        }

    actual_values = [match[1].strip() for match in matches if match[1].strip()]

    if normalized_expected_value in actual_values:
        return {
            "status": "passed",
            "details": {
                "message": f"Значение '{expected_value}' было найдено для  '{field_name}'.",
                "field": field_name,
                "expected_value": expected_value,
                "actual_value": normalized_expected_value
            }
        }

    return {
        "status": "failed",
        "details": {
            "message": f"Значение '{expected_value}' для поля '{field_name}', было найдено, но не совпало.",
            "field": field_name,
            "expected_value": expected_value,
            "actual_values": actual_values
        }
    }


def check_has_substring(file_content: str, substrings: list) -> dict:
    """ Проверяет, содержится ли хотя бы одна подстрока в файле (игнорируя регистр и пробелы) """

    normalized_content = file_content.lower().replace(" ", "").replace("\n", "")

    found_matches = []

    for substring in substrings:
        normalized_substring = substring.lower().replace(" ", "").replace("\n", "")

        if normalized_substring in normalized_content:
            found_matches.append(substring)  

    return {
        "status": "passed" if found_matches else "failed",
        "details": {
            "message": "Подстрока найдена в файле." if found_matches else "Ни одна подстрока не найдена.",
            "matches": found_matches
        }
    }


def check_no_substring(file_content: str, substrings: list) -> dict:
    """ Проверяет, что в файле нет указанных подстрок (игнорируя регистр и пробелы) """

    normalized_content = file_content.lower().replace(" ", "").replace("\n", "")

    found_matches = []

    for substring in substrings:
        normalized_substring = substring.lower().replace(" ", "").replace("\n", "")

        if normalized_substring in normalized_content:
            found_matches.append(substring)  

    return {
        "status": "failed" if found_matches else "passed",
        "details": {
            "message": "Обнаружены запрещённые подстроки!" if found_matches else "Подстроки отсутствуют.",
            "matches": found_matches
        }
    }


def not_longer_than(file_content: str, max_length: int) -> dict:
    """ Проверяет, не превышает ли количество значимых символов max_length """

    file_content = re.sub(r"//.*?$", "", file_content, flags=re.MULTILINE) 
    file_content = re.sub(r"#.*?$", "", file_content, flags=re.MULTILINE)  

    file_content = re.sub(r"/\*.*?\*/", "", file_content, flags=re.DOTALL)

    meaningful_content = re.sub(r"\s+", "", file_content)

    if len(meaningful_content) <= max_length:
        return {
            "status": "passed",
            "details": {
                "message": f"Количество значимых символов ({len(meaningful_content)}) не превышает {max_length}."
            }
        }
    else:
        return {
            "status": "failed",
            "details": {
                "message": f"Количество значимых символов ({len(meaningful_content)}) превышает {max_length}."
            }
        }


import re

def has_regex_match(file_content: str, regex_pattern: str, context_size: int = 30) -> dict:
    """ Проверяет, есть ли в файле соответствия регулярному выражению """

    regex = re.compile(regex_pattern, re.IGNORECASE)

    matches = [match for match in regex.finditer(file_content)]

    if matches:
        results = []
        for match in matches:
            start, end = match.start(), match.end()
            context_start = max(0, start - context_size)
            context_end = min(len(file_content), end + context_size)
            context = file_content[context_start:context_end]
            results.append(context.strip())

        return {
            "status": "passed",
            "details": {
                "message": f"Найдено {len(matches)} совпадение(-ий).",
                "matches": results
            }
        }
    else:
        return {
            "status": "failed",
            "details": {
                "message": "Совпадений с регулярным выражением не найдено.",
                "matches": []
            }
        }

def run_rule_check(rule_name: str, file_content: str, params: dict) -> dict:
    """ Выбирает и запускает нужную проверку в зависимости от названия правила """

    rules = {
        "check_has_string_in_file": check_has_string_in_file,
        "check_has_expected_value_in_field": check_has_expected_value_in_field,
        "check_has_substring": check_has_substring,
        "check_no_substring": check_no_substring,
        "not_longer_than": not_longer_than,
        "has_regex_match": has_regex_match
    }

    if rule_name not in rules:
        return {"status": "error", "message": f"Неизвестное правило: {rule_name}"}

    try:
        if rule_name == "check_has_string_in_file":
            return rules[rule_name](file_content, params["substring"])
        elif rule_name == "check_has_expected_value_in_field":
            return rules[rule_name](file_content, params["field_name"], params["expected_value"])
        elif rule_name in ["check_has_substring", "check_no_substring"]:
            return rules[rule_name](file_content, params["substrings"])
        elif rule_name == "not_longer_than":
            return rules[rule_name](file_content, params["max_length"])
        elif rule_name == "has_regex_match":
            return rules[rule_name](file_content, params["regex_pattern"])
        else:
            return {"status": "error", "message": "Неверные параметры для правила"}
    except KeyError as e:
        return {"status": "error", "message": f"Отсутствует параметр: {e}"}
    except Exception as e:
        return {"status": "error", "message": f"Ошибка выполнения: {str(e)}"}
    
if __name__ == "__main__":
    if len(sys.argv) != 2:
        print({"status": "error", "message": "Необходимо передать путь к файлу JSON"})
        sys.exit(1)

    file_path = sys.argv[1]

    try:
        with open(file_path, 'r', encoding='utf-8') as file:
            input_data = json.load(file)
    except FileNotFoundError:
        print({"status": "error", "message": f"Файл '{file_path}' не найден."})
        sys.exit(1)
    except json.JSONDecodeError:
        print({"status": "error", "message": "Ошибка в формате JSON в файле."})
        sys.exit(1)

    rule_name = input_data.get("rule_name")
    file_content = input_data.get("file_content")
    params = input_data.get("params")

    if not rule_name or not file_content or not isinstance(params, dict):
        print({"status": "error", "message": "Некорректные параметры JSON в файле"})
        sys.exit(1)

    result = run_rule_check(rule_name, file_content, params)

    print(json.dumps(result, ensure_ascii=False))
    
'''
Пример файла input_file.json
{
    "rule_name": "check_has_string_in_file",
    "file_content": "Hello, world!",
    "params": {
      "substring": "world"
    }
  }

Запуск скрипта 
python strict_rules.py test_file.json
'''