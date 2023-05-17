function renderVacancy (json_vacancydata) {
    let htmlCatalog = '';
    json_vacancydata.forEach(function( {id, title, salary, companyName, requirements} ) {
        htmlCatalog += `<div class="vacancy_list">
        <div class='card_top'>
        <img class="vacancy_image" src="">
        <div class='name_salary'>
        <h3>${title}</h3>
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
        <h5 class="text">О компании: ${companyName}</h5>
        <h5 class="text">Требования: ${requirements}</h5>
        </div>
        <button type="button" id="${id}" class="transform_vacancy">
            <img class="down_arrow" src="assets/images/arrow/circle_chevron.png">
        </button>
        <div class="back_side"><h5 class="text_back">Требования: ${requirements}</h5></div>
        </div>`;
        const list = document.getElementById("vacancies");
        list.innerHTML = htmlCatalog;
        renderCompetences();
    })
};

function renderCheckboxes(json_checkbox_data) {
    let htmlCatalog = '';
    let del = " слова";
    Object.keys(json_checkbox_data).forEach(function( arr ) {
        htmlCatalog += `<div id="myDropdown" class="custom-select__option">
        <label class="top_label">${arr}</label>`;
        Object.values(json_checkbox_data[arr]).forEach(function ( elem ) {
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

renderCheckboxes(FiltersJson);
renderVacancy(JSON.parse(sessionStorage.getItem('vacancies')));

