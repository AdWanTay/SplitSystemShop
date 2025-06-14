// ######################################################
// ##                    Notifications
// ######################################################

function showErr(text) {
    new Notify ({
        status: 'error',
        title: 'Ошибка',
        text: `${text}`,
        effect: 'fade',
        speed: 300,
        showIcon: true,
        showCloseButton: true,
        autoclose: true,
        autotimeout: 10000,
        type: 'outline',
        position: `right top`,
    })
};

function showNotify(title, text) {
    new Notify ({
        status: 'success',
        title: `${title}`,
        text: `${text}`,
        effect: 'fade',
        speed: 300,
        customClass: '',
        customIcon: '',
        showIcon: true,
        showCloseButton: true,
        autoclose: true,
        autotimeout: 3000,
        notificationsGap: null,
        notificationsPadding: null,
        type: 'outline',
        position: 'right top',
    })
};

function showToast(title, text) {
    new Notify ({
        status: 'info',
        title: `${title}`,
        text: `${text}`,
        effect: 'slide',
        speed: 500,
        customClass: 'customToast',
        showIcon: false,
        showCloseButton: true,
        autoclose: true,
        autotimeout: 1500,
        notificationsGap: null,
        notificationsPadding: null,
        type: 'outline',
        position: 'x-center',
    })
};


// ######################################################
// ##                    UTILITIES
// ######################################################

function setCursorPosition(pos, elem) {
    elem.focus();
    if (elem.setSelectionRange) elem.setSelectionRange(pos, pos);
    else if (elem.createTextRange) {
        const range = elem.createTextRange();
        range.collapse(true);
        range.moveEnd("character", pos);
        range.moveStart("character", pos);
        range.select()
    }
}

function mask(event) {
    // Проверяем, было ли это автозаполнение
    const isAutofill = !event.inputType && this.value.length > 0;

    let matrix = "+7 (___) ___ ____",
        i = 0,
        def = matrix.replace(/\D/g, ""),
        val = this.value.replace(/\D/g, "");

    // Обработка автозаполнения
    if (isAutofill) {
        // Если номер начинается с 7 или 8 (российский формат)
        if (/^[78]/.test(val)) {
            val = "7" + val.substring(1); // Приводим к формату +7
        }
        // Если номер начинается без кода страны (например, 900...)
        else if (val.length === 10) {
            val = "7" + val; // Добавляем 7 как код России
        }
    }

    if (def.length >= val.length) val = def;

    this.value = matrix.replace(/./g, function(a) {
        return /[_\d]/.test(a) && i < val.length ? val.charAt(i++) : i >= val.length ? "" : a;
    });

    if (event.type == "blur") {
        if (this.value.length == 2) this.value = "";
    } else {
        setCursorPosition(this.value.length, this);
    }
}
function validateEmail(email) {
    const re = /^([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22))*\x40([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d))*$/;
    return re.test(String(email).toLowerCase());
}

function lockBodyScroll() {
    const scrollbarWidth = window.innerWidth - document.documentElement.clientWidth;

    document.body.classList.add('modal-open');

    if (scrollbarWidth > 0) {
        document.body.style.paddingRight = `${scrollbarWidth}px`;
    }
}

function unlockBodyScroll() {
    document.body.classList.remove('modal-open');
    document.body.style.removeProperty('padding-right');
}

function formatPrice(price) {
    return (price / 100).toLocaleString('ru-RU'); // если цена в копейках
}

function autoResize(elem) {
    elem.style.height = 'auto';
    elem.style.height = (elem.scrollHeight-4) + 'px';
}