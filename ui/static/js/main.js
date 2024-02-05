
function showForm(formId) {
    document.getElementById(formId).style.display = 'block';
}


    function deleteArticle(articleId) {
    if (confirm('вы уверены что хотите удалить статью?')) {

    var xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/deleteArticle?id=' + articleId, true);

    xhr.onload = function () {
    if (xhr.status === 200) {
        alert('Вы успешно удалили статью');
    location.reload();
} else {

    alert('ошибка при удалении статьи, попробуйте еще раз');
}
};


    xhr.send();
}
}
function createArticle() {
    var form = $("#create-form");
    var data = form.serialize();

    $.post("/createArticle", data, function (response) {
        if (response.status === "success") {
            alert("Статья успешно создана!");
            location.reload();
        } else {
            alert("Ошибка при создании статьи!");
        }
    });
}
function updateArticle(articleId) {
    var form = $("#update-form-" + articleId);
    var data = form.serialize();

    $.post("/updateArticle", data, function (response) {
        if (response.status === "success") {
            alert("Статья успешно обновлена!");
            location.reload(); // Обновить страницу после успешного обновления
        } else {
            alert("Ошибка при обновлении статьи!");
        }
    });
}

function createDepartment() {
    var form = $("#create-form");
    var data = form.serialize();

    $.post("/createDepartment", data, function (response) {
        if (response.status === "success") {
            alert("Статья успешно создана!");
            location.reload();
        } else {
            alert("Ошибка при создании статьи!");
        }
    });
}