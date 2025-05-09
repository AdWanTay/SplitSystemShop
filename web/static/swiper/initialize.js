document.addEventListener('DOMContentLoaded', function() {
    const swiper = new Swiper('.swiper-container', {
        autoplay: {
            delay: 5000,
            disableOnInteraction: false,
        },
        loop: true,
        speed: 600, // Увеличиваем длительность анимации (в миллисекундах)
        effect: 'creative', // Используем кастомный эффект (для fade)
        creativeEffect: {
            prev: {
                opacity: 0,
                translate: [0, 0, -0.5], // Лёгкое смещение по Z (если нужно 3D-эффект)
            },
            next: {
                opacity: 0,
                translate: [0, 0, -0.5],
            },
        },
        pagination: {
            el: '.swiper-pagination',
            clickable: true,
        },
    });
});