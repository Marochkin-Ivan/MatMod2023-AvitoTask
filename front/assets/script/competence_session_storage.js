function renderCompetences () {
    let htmlCatalog = '';
    let list = document.getElementById("competence_container");
    if (JSON.parse(sessionStorage.getItem('keywords')) != null) {
    JSON.parse(sessionStorage.getItem('keywords')).forEach(function( elem ) {
        htmlCatalog += `<div class="competence">
        <div id="${elem}" class="line">
            <h7>${elem}</h7>
            <button type="button" id='cross' class="delete_competence"><img class='cross' src="assets/images/icon.png"></button>
        </div>
    </div>`;
    })
}
    if (JSON.parse(sessionStorage.getItem('employment')) != null) {
    JSON.parse(sessionStorage.getItem('employment')).forEach(function( elem ) {
        htmlCatalog += `<div class="competence">
        <div id="${elem}" class="line">
            <h7>${elem}</h7>
            <button type="button" id='cross' class="delete_competence"><img class='cross' src="assets/images/icon.png"></button>
        </div>
    </div>`;
    })
}
    if (JSON.parse(sessionStorage.getItem('schedule')) != null) {
    JSON.parse(sessionStorage.getItem('schedule')).forEach(function( elem ) {
        htmlCatalog += `<div class="competence">
        <div id="${elem}" class="line">
            <h7>${elem}</h7>
            <button type="button" id='cross' class="delete_competence"><img class='cross' src="assets/images/icon.png"></button>
        </div>
    </div>`;
    })
}  
    list.innerHTML = htmlCatalog;
    }

// function updateCheckBox () {

//     document.getElementsByClassName('checkbox').addEventListener('change', function(event) {
//         let txt = event.target.checked ? 'On' : 'Off';
//         console.log(txt);
//     });

// };


renderCompetences();

// if (document.cookie.indexOf(encodeURIComponent('user_id')) == 0) {
//     console.log('SessionStorage not cleared')
// } else {
//     sessionStorage.clear
// }

// function updateCheckBox()
// sessionStorage.getItem(key)
// sessionStorage.removeItem(key)