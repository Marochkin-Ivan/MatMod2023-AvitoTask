function first_query() {
    const query_data = document.getElementById('search').value;
    let query = query_data;
    let id_usr = decodeURIComponent(document.cookie);
    let tags = [];
    let list1 = document.querySelectorAll("#Ключевые .checkbox");
    let list2 = document.querySelectorAll("#Занятость .checkbox");
    let list3 = document.querySelectorAll("#Расписание .checkbox");
    sessionStorage.setItem('page', '1');
    let filled1 = [].filter.call( list1, function( el ) {
        return (el.checked == true) });
        filled1.forEach(el => {
          tags.push(el.getAttribute('id'))})
        console.log(tags);
        sessionStorage.setItem('keywords', JSON.stringify(tags));
        console.log(query)
        let url
    // let url = `/url?id_usr=${id_usr}&query=${query}&keywords=${(tags.join('&'))}`;
    if (tags != 0) {
      url = `http://127.0.0.1:1234/api/v1/search/search?q=${query}&keywords=${tags}`;
    } else {
      url = `http://127.0.0.1:1234/api/v1/search/search?q=${query}`;
    }
    tags = [];
    let filled2 = [].filter.call( list2, function( el ) {
      return (el.checked == true) });
      filled2.forEach(el => {
        tags.push(el.getAttribute('id'))})
      console.log(tags);
      sessionStorage.setItem('employment',JSON.stringify(tags));
    tags = [];
    let filled3 = [].filter.call( list3, function( el ) {
      return (el.checked == true) });
      filled3.forEach(el => {
        tags.push(el.getAttribute('id'))})
      console.log(tags);
      sessionStorage.setItem('schedule', JSON.stringify(tags));
      if (tags != 0) {
        url += `&schedule=${tags.join('&')}`
        }
        sessionStorage.setItem('query',url);
      url += `$page=1`
      fetch_query(url);
      setTimeout(function(){
        3
            location.reload();
            window.scrollTo(0, 0);
        4
          }, 200);
        
}

function query () {
  query = sessionStorage.getItem('query');
  page = sessionStorage.getItem('page');
  let url = `${query}$page=${page}`
  fetch_query(url);
}

document.getElementById('search').addEventListener("keypress", function(event) {
  if (event.key === "Enter") {
    event.preventDefault();
    document.querySelector(".search_button").click();
  }
});

document.querySelector(".search_button").onclick = first_query;