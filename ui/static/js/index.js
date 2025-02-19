window.onload = function() {
    let position = 0;
    const scrollBackground = () => {
        position += 1; // Увеличиваем позицию
        document.body.style.backgroundPosition = `0 -${position}px`; // Устанавливаем новую позицию фона
        if (position < 150) { // Условие для остановки (например, 100 пикселей)
            requestAnimationFrame(scrollBackground); // Запрашиваем следующую анимацию
        }
    };
    scrollBackground(); // Запускаем анимацию
};