$('#form-login').on('submit', fazerLogin)

function fazerLogin(e) {
    e.preventDefault()
   
    if($('#email').val() ===""){
        alert("Digite email...")
        $('#email').text("").focus()
        return 
    }
    if($('#senha').val() ===""){
        alert("Digite senha...")
        $('#senha').text("").focus()
        return 
    }
    
    $.ajax({
        url: 'http://localhost:3000/fazer-login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done((dados)=>{
        console.log(dados)
        console.log("deu certo")
        window.location.href = "http://localhost:3000/home";
    }).fail((erro)=>{
        console.log(erro)
        alert("Erro no Login");
    });
}