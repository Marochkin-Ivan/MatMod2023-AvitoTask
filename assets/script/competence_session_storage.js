function renderCompetences (array_competence) {
    let htmlCatalog = '';
    console.log(array_competence)
    array_competence.forEach(function( elem ) {
        htmlCatalog += `<div class="competence">
        <div id="${elem}" class="line">
            <h7>${elem}</h7>
            <button type="button" id='cross' class="delete_competence"><img class='cross' src="assets/images/icon.png"></button>
        </div>
    </div>`;
        const list = document.getElementById("competence_container");
        list.innerHTML = htmlCatalog;
    })
};

// function updateCheckBox () {

//     document.getElementsByClassName('checkbox').addEventListener('change', function(event) {
//         let txt = event.target.checked ? 'On' : 'Off';
//         console.log(txt);
//     });

// };


renderCompetences(JSON.parse(sessionStorage.getItem('keywords')));
renderCompetences(JSON.parse(sessionStorage.getItem('employment')));
renderCompetences(JSON.parse(sessionStorage.getItem('schedule')));


if (document.cookie.indexOf(encodeURIComponent('user_id')) == 0) {
    console.log('SessionStorage not cleared')
} else {
    sessionStorage.clear
}

// function updateCheckBox()
// sessionStorage.getItem(key)
// sessionStorage.removeItem(key)