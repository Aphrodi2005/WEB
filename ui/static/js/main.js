
function showForm(formId) {
    document.getElementById(formId).style.display = 'block';
}


function deleteMovie(movieId) {
    if (confirm('Вы уверены что хотите удалить статью?')) {
        var xhr = new XMLHttpRequest();
        xhr.open('DELETE', '/deleteMovie?id=' + movieId, true);

        xhr.onload = function () {
            if (xhr.status === 200) {
                alert('Вы успешно удалили статью');
                location.reload();
            } else {
                alert('Ошибка при удалении статьи, попробуйте еще раз');
            }
        };

        xhr.send();
    }
}

function createMovie() {
    var form = $("#create-form");
    var data = form.serialize();

    $.post("/createMovie", data, function (response) {
        if (response.status === "success") {
            alert("Статья успешно создана!");
            location.reload();
        } else {
            alert("Ошибка при создании статьи!");
        }
    });
}
function updateMovie(movieId) {
    var form = $("#update-form-" + movieId);
    var data = form.serialize();

    $.post("/updateMovie", data, function (response) {
        if (response.status === "success") {
            alert("Статья успешно обновлена!");
            location.reload(); // Обновить страницу после успешного обновления
        } else {
            alert("Ошибка при обновлении статьи!");
        }
    });
}

