# coding: utf-8
require 'expr_1'

RSpec.describe Expr1 do
  it 'parses test properly' do
    r = Expr1.from_file('src/str_encodings.bin')

    expect(r.len_of_1).to eq 10
    expect(r.len_of_1_mod).to eq 8
    expect(r.str1).to eq "Some ASC"
    expect(r.str1_len).to eq 8
  end
end
