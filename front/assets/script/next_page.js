function next_page(page) {
  sessionStorage.setItem("page",String(page));
  query(page);
  setTimeout(function(){
    3
      renderVacancy(JSON.parse(sessionStorage.getItem('vacancies')));
    4
      }, 50)
  setTimeout(function(){
    3
        location.reload();
        window.scrollTo(0, 0);
    4
      }, 100)
    }   

function previous_page(page) {
  console.log(page)
  sessionStorage.setItem("page",String(page));
  query(page);
  setTimeout(function(){
    3
      renderVacancy(JSON.parse(sessionStorage.getItem('vacancies')));
    4
      }, 50)
  setTimeout(function(){
    3
        location.reload();
        window.scrollTo(0, 0);
    4
      }, 100);
    }
      

page = parseInt(sessionStorage.getItem('page'));
document.querySelector(".next_page").addEventListener("click", e => {page +=1, next_page(page)});
document.querySelector(".previous_page").addEventListener("click", e => {page -=1, previous_page(page)});