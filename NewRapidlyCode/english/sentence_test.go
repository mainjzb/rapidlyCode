package english

import (
	"testing"
)

func TestSplitSentence(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "v235",
			args: args{
				"New events and major system changes are arriving this update! Become the pop star of your dreams and top the charts of popularity in Maple Pop Star Dreams! for helpful growth rewards! Spiegelmann also makes an appearance in Spiegelmann’s Starlight Box with daily rewards for growth! New cosmetics make an appearance in the return of Fairy Bros’ Golden Giveaway and special stat bonuses in Toben Hair Hero’s Special Private Course. Last but certainly not least, heaps of changes to the Better Maple experience are being added, such as the Explorer story skip function, and much more!",
			},
			want: []string{
				"New events and major system changes are arriving this update!",
				"Become the pop star of your dreams and top the charts of popularity in Maple Pop Star Dreams! for helpful growth rewards!",
				"Spiegelmann also makes an appearance in Spiegelmann’s Starlight Box with daily rewards for growth!",
				"New cosmetics make an appearance in the return of Fairy Bros’ Golden Giveaway and special stat bonuses in Toben Hair Hero’s Special Private Course.",
				"Last but certainly not least, heaps of changes to the Better Maple experience are being added, such as the Explorer story skip function, and much more!",
			},
		},
		{
			name: "v234",
			args: args{
				"The new Home offers a new way for Maplers to express their personality and interests through intricate designs! Design your dream home with furniture placement and hang out with your friends while listening to your favorite MapleStory background music! Relax in the comfort of your virtual home. When you feel like venturing outside your comfy space, take on the challenge and test your strength against the boss Kalos, fierce guardian at the top of Karote, The Unending Tower. Along with everything else, Beast Tamer and Zero character creation will be available on all worlds, including Reboot and Burning World. All of this and more await you in the Destiny: Homecoming update!",
			},
			want: []string{
				"The new Home offers a new way for Maplers to express their personality and interests through intricate designs!",
				"Design your dream home with furniture placement and hang out with your friends while listening to your favorite MapleStory background music!",
				"Relax in the comfort of your virtual home.",
				"When you feel like venturing outside your comfy space, take on the challenge and test your strength against the boss Kalos, fierce guardian at the top of Karote, The Unending Tower.",
				"Along with everything else, Beast Tamer and Zero character creation will be available on all worlds, including Reboot and Burning World.",
				"All of this and more await you in the Destiny: Homecoming update!",
			},
		},
		{
			name: "goods",
			args: args{
				"Epic Potential Scroll 50%",
			},
			want: []string{
				"Epic Potential Scroll 50%",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitParagraph(tt.args.sentence)
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("name: %v\ngot: %v\nwant: %v", tt.name, got[i], tt.want[i])
					break
				}
			}
		})
	}
}

func TestSegment(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want []PoS
	}{
		{
			name: "first Test",
			args: args{"New events and major system changes are arriving this update!"},
			want: []PoS{Adj, Noun, Conj, Adj, Noun, Noun, AuxVerb, Verb, Pron, Noun},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Segment(tt.args.sentence)
			for i := 0; i < len(got); i++ {
				if got[i].Real != tt.want[i] {
					t.Errorf("Segment() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
