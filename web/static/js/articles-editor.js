let quill;

function initializeEditor() {
    // Инициализация редактора
    quill = new Quill('#editor-container', {
        modules: {
            toolbar: [
                [{ header: [1, 2, false] }],
                ['bold', 'italic', 'underline', 'strike'],
                [{ 'list': 'ordered'}, { 'list': 'bullet' }],
                ['link', 'image']
            ]
        },
        placeholder: 'Введите текст статьи...',
        theme: 'snow'
    });

    // Обработка кнопки "Сохранить"
    document.getElementById('save-article-btn').addEventListener('click', () => {
        const title = document.getElementById('article-title').value;
        const content = quill.root.innerHTML;

        console.log('Заголовок:', title);

        console.log(quill.root.innerHTML);
        console.log(quill.getText());

        // TODO: Сделай здесь API-запрос на сохранение
        // fetch('/api/articles', { method: 'POST', body: JSON.stringify({ title, content }) ... })
    });
}