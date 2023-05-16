function like () {
    document.querySelectorAll('.transform_vacancy').forEach(element => {
        element.addEventListener('click', function() { 
            let query = this.getAttribute('id');
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
            let query = this.getAttribute('id');
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
            let query = this.getAttribute('id');
            let id_usr = decodeURIComponent(document.cookie);
            let url = `/url?id_usr=${id_usr}&id=${query}&event=respond`;
            console.log(url);
            fetch_query(url);
        })
    });
}

async function fetch_query(url) {
    let response = await fetch(url);

    if (response.ok) { // если HTTP-статус в диапазоне 200-299
         // получаем тело ответа (см. про этот метод ниже)
        let json = await response.json();
    } else {
        alert("Ошибка HTTP: " + response.status);
        }
    return json;
}

like()
write_vacancy()
respond()