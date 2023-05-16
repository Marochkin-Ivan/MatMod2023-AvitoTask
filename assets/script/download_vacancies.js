function renderVacancy (json_vacancydata) {
    let htmlCatalog = '';
    json_vacancydata.forEach(function( {id, name, salary, conditions, description, requirements} ) {
        htmlCatalog += `<div class="vacancy_list">
        <div class='card_top'>
        <img class="vacancy_image" src="">
        <div class='name_salary'>
        <h3>${name}</h3>
        <h4>${salary} ₽ </h4>
        </div>
        <button type="button" id="${id}" class="respond_vacancy">
            Откликнуться
        </button>
        <button type="button" id="${id}" class="write">
            <h5 id="write">Написать</h5>
            <img class="down_arrow" src="assets/images/icons/communication/ic_mail_outline_48px.png">
        </button>
        </div>
        <div class="description">
        <h5 class="text">Обязанности: ${conditions}</h5>
        <h5 class="text">Требования: ${requirements}</h5>
        </div>
        <button type="button" id="${id}" class="transform_vacancy">
            <img class="down_arrow" src="assets/images/arrow/circle_chevron.png">
        </button>
        <div class="back_side"><h5 class="text">О компании: ${description}</h5></div>
        </div>`;
        const list = document.getElementById("vacancies");
        list.innerHTML = htmlCatalog;
    })
};
const catalog = [
    
  {
    "id": "210000",
    "name": "ML Engineer",
    "salary": "100000",
    "experience": "3",
    "description": "Наша команда занимается разработкой системы автоматизации процессов в сфере медицины. Мы ищем специалиста, который поможет нам в разработке системы, которая будет предсказывать заболевания по симптомам.",
    "company": "ООО \"Рога и копыта\"",
    "requirements": "Опыт работы с Python от 3 лет Опыт работы с Machine Learning от 2 лет Опыт работы с Data Science от 2 лет",
    "conditions": "Офис в центре города Гибкий график работы Возможность удаленной работы"
  },
  {
    "id": "210001",
    "name": "Data Scientist",
    "salary": "80000",
    "experience": "2",
    "description": "Наша команда занимается разработкой системы автоматизации процессов в сфере медицины. Мы ищем специалиста, который поможет нам в разработке системы, которая будет предсказывать заболевания по симптомам.",
    "company": "ООО \"Рога и копыта\"",
    "requirements": "Опыт работы с Python от 3 лет Опыт работы с Machine Learning от 2 лет Опыт работы с Data Science от 2 лет",
    "conditions": "Офис в центре города Гибкий график работы Возможность удаленной работы"
  },
  {
    "id": "210002",
    "name": "Data Engineer",
    "salary": "90000",
    "experience": "2",
    "description": "Наша команда занимается разработкой системы автоматизации процессов в сфере медицины. Мы ищем специалиста, который поможет нам в разработке системы, которая будет предсказывать заболевания по симптомам.",
    "company": "ООО \"Рога и копыта\"",
    "requirements": "Опыт работы с Python от 3 лет Опыт работы с Machine Learning от 2 лет Опыт работы с Data Science от 2 лет",
    "conditions": "Офис в центре города Гибкий график работы Возможность удаленной работы"
  },
  {
    "id": "210003",
    "name": "Data Analyst",
    "salary": "70000",
    "experience": "1",
    "description": "Наша команда занимается разработкой системы автоматизации процессов в сфере медицины. Мы ищем специалиста, который поможет нам в разработке системы, которая будет предсказывать заболевания по симптомам.",
    "company": "ООО \"Рога и копыта\"",
    "requirements": "Опыт работы с Python от 3 лет Опыт работы с Machine Learning от 2 лет Опыт работы с Data Science от 2 лет",
    "conditions": "Офис в центре города Гибкий график работы Возможность удаленной работы"
  }
];

function renderCheckboxes(json_checkbox_data) {
    let htmlCatalog = '';
    let del = " слова";
    Object.keys(json_checkbox_data).forEach(function( arr ) {
        htmlCatalog += `<div id="myDropdown" class="custom-select__option">
        <label class="top_label">${arr}</label>`;
        Object.values(json_checkbox_data[arr]).forEach(function ( elem ) {
                console.log(elem)
                htmlCatalog += `<div id = "${arr.replace(del,'')}" class="extra"><label>${elem}</label>
                <input type="checkbox" class="checkbox" id="${elem}" /></div>`  
            });
        htmlCatalog += `</div>`;
        const list = document.getElementById("myDropdown");
        list.innerHTML = htmlCatalog;
    })
};

const FiltersJson = {
    "Ключевые слова": {
     "1" : "Безопасность",
     "2" : "Юриспруденция",
     "3" : "Маркетинг",
     "4" : "Консалтинг",
     "5" : "Пищевая промышленность",
     "6" : "Логистика",
     "7" : "Управление персоналом",
     "8" : "Строительство",
     "9" : "Производство",
     "10" : "Сельское хозяйство",
     "11" : "Образование",
     "12" : "Торговля",
     "13" : "Информационные технологии",
     "14" : "Электроэнергетика",
     "15" : "Здравоохранение и социальное обеспечение",
     "16" : "Бухгалтерия",
     "17" : "Реклама",
     "18" : "Туризм"},
    "Занятость": {
    "1" : "Полная занятость",
    "2" : "Частичная занятость",
    "3" : "Проектная работа",
    "4" : "Стажировка",
    "5" : "Волонтёрство"},
    "Расписание": {
    "1" : "Полный день",
    "2" : "Сменный график",
    "3" : "Гибкий график",
    "4" : "Удалённая работа",
    "5" : "Вахтовый метод"}
    }

console.log(catalog);
renderCheckboxes(FiltersJson);
// sessionStorage.setItem('vacancies',JSON.stringify(catalog))
renderVacancy(JSON.parse(sessionStorage.getItem('vacancies')));

