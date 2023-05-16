function infinite_scroll() {
  page = 1;
  while(true) {
    // нижняя граница документа
    let windowRelativeBottom = document.documentElement.getBoundingClientRect().bottom;

    // если пользователь прокрутил достаточно далеко (< 100px до конца)
    if (windowRelativeBottom < document.documentElement.clientHeight + 100) {
      // добавим больше данных
      page +=1;
      query(page);
      renderVacancy(sessionStorage.getItem('vacancies'));
    }
  }
}

infinite_scroll()