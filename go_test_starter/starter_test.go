package starter_test

import (
  "testing"
  

  "github.com/stretchr/testify/assert"
  starter "github.com/williaminfante/go_test_starter"
  "net/http/httptest"
  "io"

)

func TestSayHello_Yuri(t *testing.T) {
	greeting := starter.SayHello("Yuri")
	assert.Equal(t, "Olá, Yuri!! Seja bem vindo!!", greeting)
}

func TestSayHello_Romualdo(t *testing.T) {
	greeting := starter.SayHello("Romualdo")
	assert.Equal(t, "Olá, Romualdo!! Seja bem vindo!!", greeting)
}

// Acima, podemos perceber a prática de criar dois métodos de teste para um único método produtivo
// Isso pode ser otimizado ao utilizar o conceito de subtestes, onde nós organizamos nosso conjunto de testes para facilitar o entendimento dos nossos testes
// Abaixo, segue exemplo da aplicação de subtestes
func TestOddOrEven(t *testing.T) {
  t.Run("Apenas Positivos ou Nulo", func(t *testing.T) {
    assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
    assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
    assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
  })

  t.Run("Apenas Negativos", func(t *testing.T) {
    assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
  })

}

func TestCheckHealth(t *testing.T) {
  t.Run("Check health status", func(t *testing.T) {
    req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
    writer := httptest.NewRecorder()
    starter.CheckHealth(writer, req)
    response := writer.Result()
    body, err := io.ReadAll(response.Body)

    assert.Equal(t, "health check passed", string(body))
    assert.Equal(t, 200, response.StatusCode)
    assert.Equal(t,
                 "text/plain; charset=utf-8",
                 response.Header.Get("Content-Type"))
    assert.Equal(t, nil, err)
  })
}