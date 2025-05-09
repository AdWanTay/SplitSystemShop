// ################################################################################################################
// ##                    Notifications
// ################################################################################################################

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


document.addEventListener("DOMContentLoaded", () => {

    // ################################################################################################################
    // ##                    Profile
    // ################################################################################################################

    const profileBtn = document.getElementById('profileBtn');
    const dropdown = document.querySelector('.profile-dropdown');

    let hideTimeout;

    const showDropdown = () => {
        clearTimeout(hideTimeout);
        dropdown.classList.add('visible');
        profileBtn.classList.add('active');
    };

    const hideDropdown = () => {
        dropdown.classList.remove('visible');
        profileBtn.classList.remove('active');

        hideTimeout = setTimeout(() => {
            dropdown.style.visibility = 'hidden';
            dropdown.style.pointerEvents = 'none';
        }, 300); // такое же, как transition
    };

    const forceVisible = () => {
        dropdown.style.visibility = 'visible';
        dropdown.style.pointerEvents = 'auto';
        profileBtn.classList.add('active');
    };

    try {
        profileBtn.addEventListener('mouseenter', () => {
            forceVisible();
            showDropdown();
        });
        profileBtn.addEventListener('mouseleave', hideDropdown);
        dropdown.addEventListener('mouseenter', showDropdown);
        dropdown.addEventListener('mouseleave', hideDropdown);
    } catch { }
});


function setCursorPosition(pos, elem) {
    elem.focus();
    if (elem.setSelectionRange) elem.setSelectionRange(pos, pos);
    else if (elem.createTextRange) {
        var range = elem.createTextRange();
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
    var re = /^([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x22([^\x0d\x22\x5c\x80-\xff]|\x5c[\x00-\x7f])*\x22))*\x40([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d)(\x2e([^\x00-\x20\x22\x28\x29\x2c\x2e\x3a-\x3c\x3e\x40\x5b-\x5d\x7f-\xff]+|\x5b([^\x0d\x5b-\x5d\x80-\xff]|\x5c[\x00-\x7f])*\x5d))*$/;
    return re.test(String(email).toLowerCase());
}