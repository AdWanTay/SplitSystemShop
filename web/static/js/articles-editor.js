let quill;

function initializeEditor() {
    quill = new Quill('#editor-container', {
        modules: {
            toolbar: [
                [{ header: [1, 2, false] }],
                ['bold', 'italic', 'underline', 'strike'],
                [{ list: 'ordered' }, { list: 'bullet' }],
                ['link', 'image']
            ],
            handlers: {
                image: imageHandler // наш кастомный обработчик
            }
        },
        placeholder: 'Введите текст статьи...',
        theme: 'snow'
    });

    document.getElementById("article-image-upload").addEventListener("change", function (event) {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = function (e) {
                document.getElementById("article-image-preview").src = e.target.result;
            };
            reader.readAsDataURL(file);
        }
    });

}

function imageHandler() {
    const input = document.createElement('input');
    input.setAttribute('type', 'file');
    input.setAttribute('accept', 'image/*');
    input.click();

    input.onchange = async () => {
        const file = input.files[0];
        if (!file) return;

        const base64 = await toBase64(file);
        const range = quill.getSelection(true);

        quill.insertEmbed(range.index, 'image', base64, 'user');
    };
}

async function handlePublishClick() {
    const title = document.getElementById("article-title").value.trim();
    const description = document.getElementById("article-short-desc").value.trim();
    const content = quill.root.innerHTML;

    const imageFile = document.getElementById("article-image-upload").files[0];
    let imageBase64 = null;
    if (imageFile) {
        imageBase64 = await toBase64(imageFile);
    }

    const articleData = {
        title,
        description,
        content,      // здесь уже встроены base64-изображения
        imageBase64   // превью
    };

    try {
        const res = await fetch("/api/articles", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(articleData)
        });

        if (!res.ok) {
            const err = await res.json();
            showErr(err.error || "Ошибка при создании статьи");
            return;
        }

        showNotify("Успех", "Статья создана");
        closeAllModals();
        location.reload();
    } catch (e) {
        showErr("Ошибка при создании: " + e.message);
    }
}

async function getArticle(id) {
    try {
        const res = await fetch(`/api/articles/${id}`);
        if (!res.ok) {
            const err = await res.json();
            throw new Error(err.error || "Ошибка при получении статьи");
        }

        const article = await res.json();

        return {
            id: article.id,
            title: article.title,
            description: article.description,
            content: article.content,
            imageURL: article.image_url,
        };
    } catch (err) {
        showErr("Не удалось загрузить статью: " + err.message);
        return null;
    }
}
async function updateArticle(id) {
    const title = document.getElementById("article-title").value.trim();
    const description = document.getElementById("article-short-desc").value.trim();
    const content = quill.root.innerHTML;
    const imageFile = document.getElementById("article-image-upload").files[0];

    let imageBase64 = null;
    if (imageFile) {
        imageBase64 = await toBase64(imageFile);
    }

    const articleData = {
        title,
        description,
        content,
        imageBase64
    };

    try {
        const res = await fetch(`/api/articles/${id}`, {
            method: "PATCH",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(articleData)
        });

        if (!res.ok) {
            const err = await res.json();
            showErr(err.error || "Ошибка при обновлении");
            return;
        }

        const updated = await res.json();

        // Обновляем карточку в DOM
        const card = document.querySelector(`.blog-card[data-id="${id}"]`);
        if (card) {
            card.querySelector(".blog-card-title").textContent = updated.title;
            card.querySelector(".blog-card-description").textContent = updated.description;
            if (updated.image_url && updated.image_url.trim() !== "") {
                card.querySelector(".blog-card-action").src = updated.image_url;
            }
        }

        showNotify("Успех", "Статья обновлена");
        closeAllModals();
    } catch (e) {
        showErr("Ошибка при обновлении: " + e.message);
    }
}


function toBase64(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => resolve(reader.result);
        reader.onerror = (err) => reject(err);
        reader.readAsDataURL(file);
    });
}
