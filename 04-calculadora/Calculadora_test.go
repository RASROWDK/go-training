package main

import "testing"

func Test_adicionarPercentual(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_descobrirPercentual(t *testing.T) {
	type args struct {
		valor  float64
		valor2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := descobrirPercentual(tt.args.valor, tt.args.valor2); got != tt.want {
				t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplicaçao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplicaçao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiplicaçao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{

		// TODO: Add test cases.
		{
			name: "test 1+1",
			args: args{1, 1},
			want: 2,
		},
       {

			name: "test 1+2",
	        args: args{1,2},
			want: 3,

       },
       {
			name: "test 10 + 10",
			args: args{10,10},
            want:20,
       },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtraçao(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtraçao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("subtraçao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pow(t *testing.T) {
	type args struct {
		valor    float64
		multiplo float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	{
		name:"10 potencia 2",
		args: args{10,3},
	    want:1000,
	},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pow(tt.args.valor, tt.args.multiplo); got != tt.want {
				t.Errorf("pow() = %v, want %v", got, tt.want)
			}
		})
	}
}