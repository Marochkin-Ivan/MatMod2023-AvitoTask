function first_query() {
    const query_data = document.getElementById('search').value;
    // let xhr = new XMLHttpRequest();
    let query = query_data;
    let id_usr = decodeURIComponent(document.cookie);
    let tags = [];
    let list1 = document.querySelectorAll("#Ключевые .checkbox");
    let list2 = document.querySelectorAll("#Занятость .checkbox");
    let list3 = document.querySelectorAll("#Расписание .checkbox");
    let filled1 = [].filter.call( list1, function( el ) {
        return (el.checked == true) });
        filled1.forEach(el => {
          tags.push(el.getAttribute('id'))})
        console.log(tags);
        sessionStorage.setItem('keywords', JSON.stringify(tags));
    // let url = `/url?id_usr=${id_usr}&query=${query}&keywords=${(tags.join('&'))}`;
    let url = `http://127.0.0.1:1234/api/v1/search/search?q=${query}&keywords=${tags}`;
    tags = [];
    let filled2 = [].filter.call( list2, function( el ) {
      return (el.checked == true) });
      filled2.forEach(el => {
        tags.push(el.getAttribute('id'))})
      console.log(tags);
      sessionStorage.setItem('employment',JSON.stringify(tags));
    url += `&employment=${tags}`;
    tags = [];
    let filled3 = [].filter.call( list3, function( el ) {
      return (el.checked == true) });
      filled3.forEach(el => {
        tags.push(el.getAttribute('id'))})
      console.log(tags);
      sessionStorage.setItem('schedule', JSON.stringify(tags));
    // url += `&schedule=${tags.join('&')}$page=1`
    vacancy_json = fetch_query(url);
    sessionStorage.setItem('vacancies',JSON.stringify(vacancy_json));
    return vacancy_json
        }

function query (page) {
  const query_data = document.getElementById('search').value;
  // let xhr = new XMLHttpRequest();
  let query = query_data;
  let id_usr = decodeURIComponent(document.cookie);
  let tags = [];
  let list1 = document.querySelectorAll("#Ключевые .checkbox");
  let list2 = document.querySelectorAll("#Занятость .checkbox");
  let list3 = document.querySelectorAll("#Расписание .checkbox");
  let filled1 = [].filter.call( list1, function( el ) {
      return (el.checked == true) });
      filled1.forEach(el => {
        tags.push(el.getAttribute('id'))})
      console.log(tags);
  // let url = `/url?id_usr=${id_usr}&query=${query}&keywords=${(tags.join('&'))}`;
  let url = `http://127.0.0.1:1234/api/v1/search/search?q=${query}&`;
  tags = [];
  let filled2 = [].filter.call( list2, function( el ) {
    return (el.checked == true) });
    filled2.forEach(el => {
      tags.push(el.getAttribute('id'))})
    console.log(tags);
  url += `&employment=${tags}`;
  tags = [];
  let filled3 = [].filter.call( list3, function( el ) {
    return (el.checked == true) });
    filled3.forEach(el => {
      tags.push(el.getAttribute('id'))})
    console.log(tags);
  url += `&schedule=${tags.join('&')}$page=${page}`
  vacancy_json = fetch_query(url);
  sessionStorage.setItem('vacancies',JSON.stringify(vacancy_json));
  return vacancy_json
      }

    // xhr.open("GET", '/submit?' + params, true);
    // xhr.send();


document.getElementById('search').addEventListener("keypress", function(event) {
    if (event.key === "Enter") {
      event.preventDefault();
      document.querySelector(".search_button").click();
    }
  });

document.querySelector(".search_button").onclick = first_query;
