document.addEventListener('DOMContentLoaded', function() {
    // Sidebar navigation
    document.querySelectorAll('.sidebar .nav-item').forEach(item => {
        item.addEventListener('click', function() {
            // Remove active class from all items
            document.querySelectorAll('.sidebar .nav-item').forEach(i => i.classList.remove('active'));
            
            // Add active class to the clicked item
            this.classList.add('active');
        });
    });

    // Signup button
    const signupButton = document.getElementById('join-community-btn');
    if (signupButton) {
        signupButton.addEventListener('click', function() {
            window.location.href = '/register';
        });
    }
});