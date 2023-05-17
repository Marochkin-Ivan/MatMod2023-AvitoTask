function like () {
    document.querySelectorAll('.transform_vacancy').forEach(element => {
        element.addEventListener('click', function() { 
            let query = element.getAttribute('id');
            let id_usr = decodeURIComponent(document.cookie);
            let url = `/url?id_usr=${id_usr}&id_vac=${query}&event=like`;
            console.log(url)
            fetch_query(url);
        })
    });
}
    

function write_vacancy () {
    document.querySelectorAll('.write').forEach(element => {
        element.addEventListener('click', function() { 
            let query = element.getAttribute('id');
            let id_usr = decodeURIComponent(document.cookie);
            let url = `/url?id_usr=${id_usr}&id=${query}&event=write`;
            console.log(url)
            fetch_query(url);
        })
    });
}

function respond () {
    document.querySelectorAll('.respond_vacancy').forEach(element => {
        element.addEventListener('click', function() { 
            let query = element.getAttribute('id');
            let id_usr = decodeURIComponent(document.cookie);
            let url = `/url?id_usr=${id_usr}&id=${query}&event=respond`;
            console.log(url);
            fetch_query(url);
        })
    });
}

async function fetch_query(url) {
    sessionStorage.removeItem('vacancies');
    let response = await fetch(url);
    console.log(url)
    if (response.ok) { // если HTTP-статус в диапазоне 200-299
         // получаем тело ответа (см. про этот метод ниже)
        let json_answer = await response.json();
        console.log(json_answer)
        sessionStorage.setItem('vacancies',JSON.stringify(json_answer));
    } else {
        alert("Ошибка HTTP: " + response.status);
        }
}

like()
write_vacancy()
respond()