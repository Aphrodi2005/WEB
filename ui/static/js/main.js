$(document).ready(function () {
    $(".update-article-btn").on("click", function () {
        var form = $(this).closest("form");
        var data = form.serialize();

        // Получить значения нового заголовка, контента и категории
        var newTitle = $("#new-title").val();
        var newContent = $("#new-content").val();
        var newCategory = $("#new-category").val();

        // Добавить их к данным формы
        data += "&new-title=" + encodeURIComponent(newTitle);
        data += "&new-content=" + encodeURIComponent(newContent);
        data += "&new-category=" + encodeURIComponent(newCategory);

        $.post("/updateArticle", data, function (response) {
            if (response.status === "success") {
                alert("Статья успешно обновлена!");
            } else {
                alert("Ошибка при обновлении статьи!");
            }
        });
    });



    $(".delete-article-btn").on("click", function () {
        var form = $(this).closest("form");
        var data = form.serialize();

        $.post("/deleteArticle", data, function (response) {
            if (response.status === "success") {
                // Удаляем контейнер статьи (здесь вам нужно удалить весь блок статьи)
                alert("Статья успешно удалена!");
            } else {
                alert("Ошибка при удалении статьи!");
            }
        });
    });

});

    function deleteArticle(articleId) {
    if (confirm('Are you sure you want to delete this article?')) {
    // Создайте XMLHttpRequest-запрос
    var xhr = new XMLHttpRequest();
    xhr.open('DELETE', '/deleteArticle?id=' + articleId, true);

    // Установите обработчик события для обработки ответа
    xhr.onload = function () {
    if (xhr.status === 200) {
    // Обновите страницу после успешного удаления
    location.reload();
} else {
    // Обработайте ошибку при удалении
    alert('Failed to delete the article. Please try again.');
}
};

    // Отправьте запрос
    xhr.send();
}
}
