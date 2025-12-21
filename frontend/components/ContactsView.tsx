import React, { useState } from 'react';
import { INTERNAL_CONTACTS } from '../constants';
import { Phone, Mail, Search, Users, Building, Hash } from 'lucide-react';

const ContactsView: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState('');

  const filteredContacts = INTERNAL_CONTACTS.filter(contact => 
    contact.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    contact.department.toLowerCase().includes(searchTerm.toLowerCase()) ||
    contact.role.toLowerCase().includes(searchTerm.toLowerCase()) ||
    contact.extension.includes(searchTerm)
  );

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      
      {/* Header & Search */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
            <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-3">
                <Phone className="text-emerald-500" />
                Ramais e E-mails
            </h1>
            <p className="text-gray-500">Lista de contatos internos da Solis</p>
        </div>
        
        <div className="relative w-full md:w-72">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" size={18} />
            <input 
              type="text" 
              placeholder="Buscar por nome, setor ou ramal..." 
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="w-full pl-10 pr-4 py-2 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm placeholder-gray-600"
            />
        </div>
      </div>

      {/* Table */}
      <div className="bg-[#1a1d24] rounded-xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
        <div className="overflow-x-auto custom-scrollbar flex-grow">
            <table className="w-full text-sm text-left">
                <thead className="bg-[#20232b] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0 z-10">
                    <tr>
                        <th className="px-6 py-4">Nome</th>
                        <th className="px-6 py-4">Cargo</th>
                        <th className="px-6 py-4">Departamento</th>
                        <th className="px-6 py-4 text-center">Ramal</th>
                        <th className="px-6 py-4">E-mail</th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-gray-800">
                    {filteredContacts.length === 0 ? (
                        <tr>
                            <td colSpan={5} className="px-6 py-12 text-center text-gray-500">
                                Nenhum contato encontrado para "{searchTerm}".
                            </td>
                        </tr>
                    ) : (
                        filteredContacts.map((contact, index) => (
                            <tr key={index} className="hover:bg-[#20232b] transition-colors group">
                                <td className="px-6 py-4">
                                    <div className="flex items-center gap-3">
                                        <div className="w-8 h-8 rounded-full bg-gray-800 flex items-center justify-center text-gray-400 font-bold border border-gray-700">
                                            {contact.name.charAt(0)}
                                        </div>
                                        <span className="font-medium text-gray-200">{contact.name}</span>
                                    </div>
                                </td>
                                <td className="px-6 py-4 text-gray-400">
                                    {contact.role || '-'}
                                </td>
                                <td className="px-6 py-4">
                                    <div className="flex items-center gap-2 text-gray-300">
                                        {contact.department && <Building size={14} className="text-gray-600" />}
                                        {contact.department || '-'}
                                    </div>
                                </td>
                                <td className="px-6 py-4 text-center">
                                    {contact.extension ? (
                                        <span className="inline-flex items-center justify-center px-3 py-1 bg-[#1e5144]/10 text-emerald-400 rounded-md font-mono font-bold border border-[#1e5144]/20">
                                            {contact.extension}
                                        </span>
                                    ) : '-'}
                                </td>
                                <td className="px-6 py-4">
                                    {contact.email ? (
                                        <div className="flex items-center gap-2 text-blue-400 hover:text-blue-300 transition-colors cursor-pointer" onClick={() => window.location.href = `mailto:${contact.email}`}>
                                            <Mail size={14} />
                                            {contact.email}
                                        </div>
                                    ) : '-'}
                                </td>
                            </tr>
                        ))
                    )}
                </tbody>
            </table>
        </div>
        <div className="p-4 border-t border-gray-800 bg-[#15171c] text-xs text-gray-500 text-center">
            Total de {filteredContacts.length} contatos listados.
        </div>
      </div>
    </div>
  );
};

export default ContactsView;