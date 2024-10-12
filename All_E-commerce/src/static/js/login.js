document.getElementById('loginForm').addEventListener('submit', async function(event) {
    event.preventDefault(); // Empêche le comportement par défaut du formulaire
    const formData = new FormData(event.target); // Récupère les données du formulaire
    const data = {
        email: formData.get('email'),
        password: formData.get('password')
    };
    
    const response = await fetch('/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    });
    
    if (response.ok) {
        const result = await response.json();
        localStorage.setItem("token", result.token); // Stocke le token dans le localStorage
        alert("Connexion réussie");
        // Redirige l'utilisateur vers une autre page si nécessaire
        window.location.href = "/dashboard"; // Changez cela selon vos besoins
    } else {
        alert("Erreur lors de la connexion");
    }
});
