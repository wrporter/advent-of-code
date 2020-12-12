package main

import (
	"fmt"
	"testing"
)

func Test_winMinMana(t *testing.T) {
	type args struct {
		startPlayer Character
		startBoss   Character
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				startPlayer: Character{
					HitPoints: 10,
					Mana:      250,
				},
				startBoss: Character{
					HitPoints: 13,
					Damage:    9,
				},
			},
			24,
		},
		{
			args{
				startPlayer: Character{
					HitPoints: 10,
					Mana:      250,
				},
				startBoss: Character{
					HitPoints: 14,
					Damage:    9,
				},
			},
			114,
		},
		//{
		//	args{
		//		startPlayer: Character{
		//			HitPoints: 50,
		//			Mana:      500,
		//		},
		//		startBoss: Character{
		//			HitPoints: 71,
		//			Damage:    10,
		//		},
		//	},
		//	114,
		//},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			if got := winMinMana(tt.args.startPlayer, tt.args.startBoss, ""); got != tt.want {
				t.Errorf("winMinMana() = %v, want %v", got, tt.want)
			}
		})
	}
}
