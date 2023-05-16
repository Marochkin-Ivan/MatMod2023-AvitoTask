const $right_button = document.getElementById('right_arrow');
const $left_button = document.getElementById('left_arrow');
const $form1 = document.querySelector('#competence_container').lastElementChild;
const $form2 = document.querySelector('#competence_container').firstElementChild;


$right_button.addEventListener('click', e => {
  // Прокрутим страницу к форме 
  $form1.scrollIntoView({ 
    block: 'end', // к ближайшей границе экрана
    behavior: 'smooth', // и плавно 
  });
});

$left_button.addEventListener('click', e => {
    // Прокрутим страницу к форме 
    $form2.scrollIntoView({ 
      block: 'end', // к ближайшей границе экрана
      behavior: 'smooth', // и плавно 
    });
  });

// // Удаление компетенции
// function remove(){
//     let w = this.closest('.competence'), br = w.nextElementSibling;
    
//     if(br.nodeType === 1 && br.nodeName === 'BR')
//       br.remove();
    
//     w.remove();
//   }
  
//   Array.from(document.querySelectorAll('.cross')).forEach(b => b.addEventListener('click', remove));


//Запрос Id пользователя и сохранение их в cookie
if (document.cookie.indexOf(encodeURIComponent('user_id')) == 0) {
    let delete_id = confirm("Dev_хотите удалить Id пользователя?");
    if (delete_id == true) {
        setCookie(user_id, "", {
            'max-age': -1
          })
    }else {
    alert(decodeURIComponent(document.cookie));
    }
    
}else{
    let user_id = prompt("Dev_введите Id для текущего пользователя:", undefined);
    document.cookie = encodeURIComponent('user_id') + '=' + encodeURIComponent(user_id);
}