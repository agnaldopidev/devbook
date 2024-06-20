$('#form-cadastro').on('submit', cadastro)

function cadastro(e) {
    e.preventDefault()
    //console.log("Cadastro44")
    //console.log($('#nome').val())

    if($('#senha').val() != $('#confirmarSenha').val()){
        alert("Senha deve igual confimar senha...")
    }

    if($('#nome').val() ===''){
        alert("Digite nome...")
        $('#nome').text("").focus()
        return
    }
    if($('#nick').val() ===""){
        alert("Digie nick...")
        $('#nick').text("").focus()
        return
    }
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
        url: 'http://localhost:3000/cadastro',
        method: 'POST',
        data: {
            nome: $('#nome').val(),
            nick: $('#nick').val(), 
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(()=>{
        alert("Cadastro realizado com sucesso");
        $('#form-cadastro')[0].reset();
    }).fail((erro)=>{
        console.log(erro)
        alert("Erro no cadastro");
    });
}