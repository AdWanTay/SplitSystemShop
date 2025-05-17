// blog.js - скрипты для работы над статьями

function createArticleEditor() {
    closeAllModals();
    lockBodyScroll();

    fetch("/web/templates/components/modals/article-create-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            initializeEditor();

            // Навешиваем обработчик только после вставки в DOM
            const publishBtn = document.getElementById("publish-article-btn");
            if (publishBtn) {
                publishBtn.addEventListener("click", handlePublishClick);
            }
        });
}


async function openArticleEditor(id) {
    closeAllModals();
    lockBodyScroll();

    const html = await fetch("/web/templates/components/modals/article-edit-modal.html").then(res => res.text());
    const modalContainer = document.createElement("div");
    modalContainer.innerHTML = html;
    document.body.appendChild(modalContainer);

    initializeEditor();

    // Загрузим данные статьи после инициализации редактора
    const article = await getArticle(id);
    if (!article) return;

    document.getElementById("article-title").value = article.title;
    document.getElementById("article-short-desc").value = article.description;
    quill.root.innerHTML = article.content;
    const imageURL = article.imageURL;

    document.getElementById("article-image-preview").src = imageURL;
    document.getElementById("save-article-btn").addEventListener("click", () => updateArticle(id));
}

async function articleDeleteConfirm(id) {
    const article = await getArticle(id);
    const article_title = article?.title || "Нет названия";

    openModal({
        title: "Подтверждение удаления",
        body: `<p>Вы действительно хотите удалить статью: "${article_title}"?</p>`,
        description: "ⓘ Это действие необратимо, статья будет удалено безвозвратно",
        mainBtnText: "Удалить статью",
        mainBtnAction: async function () {
            try {
                const res = await fetch(`/api/articles/${id}`, {
                    method: 'DELETE',
                });

                if (!res.ok) {
                    const data = await res.json();
                    showErr(data.error || "Не удалось удалить статью");
                    return;
                }

                showNotify("Успех", "Статья успешно удалена");
                closeAllModals();
                
                // Удаляем элемент статьи из DOM
                const articleCard = document.querySelector(`.blog-card[data-id="${id}"]`);
                if (articleCard) {
                    articleCard.remove();
                }
            } catch (err) {
                showErr("Ошибка при удалении: " + err.message);
            }
        }
    });
}