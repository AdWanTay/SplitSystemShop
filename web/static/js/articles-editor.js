let quill;
let imageBase64 = null;

function initializeEditor() {
    // Инициализация редактора
    quill = new Quill('#editor-container', {
        modules: {
            toolbar: [
                [{ header: [1, 2, false] }],
                ['bold', 'italic', 'underline', 'strike'],
                [{ list: 'ordered' }, { list: 'bullet' }],
                ['link', 'image']
            ]
        },
        placeholder: 'Введите текст статьи...',
        theme: 'snow'
    });

    // Обработка загрузки картинки
    document.getElementById('article-image-upload').addEventListener('change', function () {
        const file = this.files[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onload = function (e) {
            imageBase64 = e.target.result;
            document.getElementById('article-image-preview').src = imageBase64;
        };
        reader.readAsDataURL(file);
    });

    // Обработка кнопки "Сохранить"
    document.getElementById('save-article-btn').addEventListener('click', async () => {
        const title = document.getElementById('article-title').value.trim();
        const description = document.getElementById('article-short-desc').value.trim();
        const content = quill.root.innerHTML.trim();

        if (!title || !description || !content || !imageBase64) {
            showErr('Заполните все поля и загрузите картинку');
            return;
        }

        try {
            const response = await fetch('/api/articles', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    title,
                    description,
                    content,
                    image_url: imageBase64
                })
            });

            const result = await response.json();

            if (response.ok) {
                closeAllModals(); // или document.getElementById('articleEditorModal').style.display = 'none';
                location.reload();
            } else {
                showErr('Ошибка: ' + (result.error || 'Неизвестная ошибка'));
            }
        } catch (error) {
            console.error('Ошибка:', error);
            showErr('Ошибка при отправке запроса');
        }
    });
}
