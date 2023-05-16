var page = 1;
var pageSize = 20;
var searchQuery = "";

function renderVacancyCard(vacancy) {
    var card = document.createElement("div");
    card.className = "vacancy-card";

    var title = document.createElement("h2");
    title.textContent = vacancy.title;
    card.appendChild(title);

    var description = document.createElement("p");
    description.textContent = vacancy.description;
    card.appendChild(description);

    var salary = document.createElement("p");
    salary.textContent = vacancy.salary;
    card.appendChild(salary);

    var tags = document.createElement("div");
    tags.className = "tag-container";
    vacancy.tags.forEach(function(tag) {
        var tagElement = document.createElement("span");
        tagElement.className = "tag";
        tagElement.textContent = tag;
        tags.appendChild(tagElement);
    });
    card.appendChild(tags);

    document.getElementById("vacancy-list").appendChild(card);
}

// Базовая карточка вакансии для отображения по умолчанию
var defaultVacancy = {
    title: "Название вакансии",
    description: "Описание вакансии",
    salary: "Заработная плата",
    tags: ["Тег1", "Тег2", "Тег3"]
};

renderVacancyCard(defaultVacancy);


function loadMoreResults() {
    page++; // Увеличиваем номер страницы

    var url = "http://localhost:9999/api/v1/search?query=" + encodeURIComponent(searchQuery) + "&page=" + page + "&pageSize=" + pageSize;

    // Здесь вы можете выполнить AJAX-запрос или использовать другие методы отправки запроса на сервер
    // Например, с использованием Fetch API или XMLHttpRequest

    // Пример использования Fetch API
    fetch(url)
        .then(function(response) {
            return response.json(); // Преобразуем ответ в JSON
        })
        .then(function(data) {
            // Обработка полученных данных
            document.getElementById("vacancy-list").innerHTML = ""; // Очищаем список вакансий

            data.forEach(function(vacancy) {
                renderVacancyCard(vacancy); // Отображаем каждую вакансию
            });

            if (data.length < pageSize) {
                // Если получено меньше результатов, чем pageSize, скрываем кнопку "Загрузить еще"
                document.getElementById("load-more-container").style.display = "none";
            } else {
                // Если есть еще результаты, показываем кнопку "Загрузить еще"
                document.getElementById("load-more-container").style.display = "block";
            }
        })
        .catch(function(error) {
            // Обработка ошибок
        });
}

document.getElementById("load-more-button").addEventListener("click", function() {
    loadMoreResults();
});